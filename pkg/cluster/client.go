package cluster

import (
	"context"
	"fmt"
	"sync"
	"time"

	"capnproto.org/go/capnp/v3"
	"github.com/libp2p/go-libp2p/core/peer"

	api "github.com/wetware/casm/internal/api/routing"
	casm "github.com/wetware/casm/pkg"
	"github.com/wetware/casm/pkg/cluster/pulse"
	"github.com/wetware/casm/pkg/cluster/routing"
)

type View api.View

func (v View) AddRef() View {
	return View(capnp.Client(v).AddRef())
}

func (v View) Release() {
	capnp.Client(v).Release()
}

// Lookup returns the first record to match the supplied query.
// Callers MUST call ReleaseFunc when finished. Note that this
// will invalidate any record returned by FutureRecord.
func (v View) Lookup(ctx context.Context, query Query) (FutureRecord, capnp.ReleaseFunc) {
	f, release := api.View(v).Lookup(ctx, func(ps api.View_lookup_Params) error {
		return query(ps)
	})

	return FutureRecord(f.Result()), release
}

// Iter returns an iterator that ranges over records matching
// the supplied query. Callers MUST call the ReleaseFunc when
// finished with the iterator.  Callers MUST NOT call methods
// on the iterator after calling the ReleaseFunc.
func (v View) Iter(ctx context.Context, query Query) (*Iterator, capnp.ReleaseFunc) {
	var (
		h          = make(handler)
		f, release = api.View(v).Iter(ctx, h.Handler(query))
	)

	return &Iterator{
		f: casm.Future(f),
		h: h,
	}, release
}

// Iterator is a stateful object that enumerates routing
// records.
//
// Callers SHOULD check the value of Err after a call to
// Next returns nil.  If Err() == nil and Next() == nil,
// the iterator is exhausted.
type Iterator struct {
	f    casm.Future
	h    handler
	curr *borrowedRecord
}

// Err returns any error encountered by the iterator.
// If Err() != nil, future calls to Err() return the
// same error, and calls to Next() return nil.
//
// Callers SHOULD check Err() after a call to Next()
// returns nil.
func (it *Iterator) Err() error {
	select {
	case <-it.f.Done():
		return it.f.Err()
	default:
		return nil
	}
}

// Next upates the iterator's internal state and returns the
// next record in the stream.  If a call to Next returns nil,
// the iterator is exhausted.
//
// Records returned by Next are valid until the next call to
// Next, or until the iterator is released.  See View.Iter().
func (it *Iterator) Next() routing.Record {
	it.curr.Release()

	it.curr = <-it.h        // closed when handler is released
	return it.curr.Record() // nil when it.h is closed
}

type borrowedRecordPool sync.Pool

var pool borrowedRecordPool

func (p *borrowedRecordPool) Borrow(r routing.Record) *borrowedRecord {
	if v := (*sync.Pool)(p).Get(); v != nil {
		br := v.(*borrowedRecord)
		br.rec = r
		return br
	}

	return &borrowedRecord{
		rec:  r,
		done: make(chan struct{}, 1),
	}
}

func (p *borrowedRecordPool) Return(r *borrowedRecord) {
	(*sync.Pool)(p).Put(r)
}

type borrowedRecord struct {
	rec  routing.Record
	done chan struct{}
}

func (br *borrowedRecord) Done() <-chan struct{} {
	return br.done
}

func (br *borrowedRecord) Release() {
	if br != nil {
		br.done <- struct{}{}
	}
}

func (br *borrowedRecord) Record() (r routing.Record) {
	if br != nil {
		r = br.rec
	}

	return
}

type handler chan *borrowedRecord

func (ch handler) Shutdown() { close(ch) }

func (ch handler) Handler(query Query) func(api.View_iter_Params) error {
	return func(ps api.View_iter_Params) (err error) {
		if err = query(ps); err == nil {
			err = ps.SetHandler(api.View_Handler_ServerToClient(ch))
		}

		return
	}
}

func (ch handler) Recv(ctx context.Context, call api.View_Handler_recv) error {
	rec, err := call.Args().Record()
	if err != nil {
		return err
	}

	r, err := newRecord(rec)
	if err != nil {
		return err
	}

	borrowed := pool.Borrow(r)
	defer pool.Return(borrowed)

	select {
	case ch <- borrowed:
	case <-ctx.Done():
		return ctx.Err()
	}

	select {
	case <-borrowed.Done():
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

type FutureRecord api.View_MaybeRecord_Future

func (f FutureRecord) Await(ctx context.Context) (routing.Record, error) {
	select {
	case <-f.Done():
		return f.Record()

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (f FutureRecord) Record() (routing.Record, error) {
	res, err := api.View_MaybeRecord_Future(f).Struct()
	if err != nil {
		return nil, err
	}

	if !res.HasJust() {
		return nil, nil // no record
	}

	rec, err := res.Just()
	if err != nil {
		return nil, err
	}

	return newRecord(rec)
}

type clientRecord api.View_Record

func newRecord(rec api.View_Record) (clientRecord, error) {
	// validate record fields

	hb, err := rec.Heartbeat()
	if err != nil {
		return clientRecord{}, fmt.Errorf("heartbeat: %w", err)
	}

	if _, err := hb.Meta(); err != nil {
		return clientRecord{}, fmt.Errorf("meta: %w", err)
	}

	// use FooBytes to avoid allocating a string
	if _, err := hb.HostBytes(); err != nil {
		return clientRecord{}, fmt.Errorf("host: %w", err)
	}

	if _, err := rec.PeerBytes(); err != nil {
		return clientRecord{}, fmt.Errorf("peer:  %w", err)
	}

	return clientRecord(rec), nil
}

func (r clientRecord) Peer() peer.ID {
	id, _ := api.View_Record(r).Peer()
	return peer.ID(id)
}

func (r clientRecord) PeerBytes() ([]byte, error) {
	return api.View_Record(r).PeerBytes()
}

func (r clientRecord) Seq() uint64 {
	return api.View_Record(r).Seq()
}

func (r clientRecord) TTL() time.Duration {
	return r.heartbeat().TTL()
}

func (r clientRecord) Instance() uint32 {
	return r.heartbeat().Instance()
}

func (r clientRecord) Host() (string, error) {
	return r.heartbeat().Host()
}

func (r clientRecord) HostBytes() ([]byte, error) {
	return r.heartbeat().HostBytes()
}

func (r clientRecord) Meta() (routing.Meta, error) {
	return r.heartbeat().Meta()
}

func (r clientRecord) heartbeat() pulse.Heartbeat {
	hb, _ := api.View_Record(r).Heartbeat()
	return pulse.Heartbeat{Heartbeat: hb}
}

func (r clientRecord) BindRecord(rec api.View_Record) (err error) {
	if err = rec.SetPeer(string(r.Peer())); err == nil {
		rec.SetSeq(r.Seq())
		err = rec.SetHeartbeat(r.heartbeat().Heartbeat)
	}

	return
}
