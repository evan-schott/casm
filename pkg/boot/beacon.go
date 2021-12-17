package boot

import (
	"bytes"
	"context"
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/host"
	ps "github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/lthibault/log"
)

// Beacon is a small discovery server that binds to a local address
// and replies to incoming connections with the Host's peer record.
type Beacon struct {
	Logger log.Logger
	Addr   net.Addr
	Host   host.Host
}

// Used by supervisor to report the service name on failure.
func (b Beacon) String() string {
	return "casm.boot.beacon"
}

func (b Beacon) Advertise(ctx context.Context, ns string, opt ...discovery.Option) (time.Duration, error) {
	b.Logger.WithField("ttl", ps.PermanentAddrTTL).
		Warn("stub call to advertise returned")

	return ps.PermanentAddrTTL, nil
}

func (b Beacon) Serve(ctx context.Context) error {
	sub, err := b.Host.EventBus().Subscribe(new(event.EvtLocalAddressesUpdated))
	if err != nil {
		return err
	}
	defer sub.Close()

	server, err := new(net.ListenConfig).Listen(ctx,
		b.Addr.Network(),
		b.Addr.String())
	if err != nil {
		return err
	}
	defer server.Close()

	payload, err := newAtomicPayload(ctx, sub)
	if err != nil {
		return err
	}

	requests := make(chan net.Conn, 1)
	defer close(requests)

	go func() {
		for v := range sub.Out() {
			payload.ConsumeEvent(v.(event.EvtLocalAddressesUpdated))
		}
	}()

	for ctx.Err() == nil {
		conn, err := server.Accept()
		if err != nil {
			return err
		}

		go func(conn net.Conn) {
			defer conn.Close()

			err := conn.SetWriteDeadline(time.Now().Add(time.Second))
			if err != nil {
				b.Logger.WithError(err).Debug("failed to set deadline")
				return
			}

			n, err := payload.WriteTo(conn)
			if err != nil {
				b.Logger.WithError(err).Debug("failed to write payload")
				return
			}

			b.Logger.WithField("bytes", n).Debug("wrote payload")

		}(conn)

	}

	return ctx.Err()
}

type atomicPayload atomic.Value

func newAtomicPayload(ctx context.Context, sub event.Subscription) (*atomicPayload, error) {
	var ap atomicPayload
	select {
	case v := <-sub.Out():
		err := ap.ConsumeEvent(v.(event.EvtLocalAddressesUpdated))
		return &ap, err

	case <-ctx.Done():
		// This usually occurs because the host isn't listening on any addresses.
		return nil, ctx.Err()
	}
}

func (ap *atomicPayload) WriteTo(w io.Writer) (int64, error) {
	return io.Copy(w, bytes.NewReader(ap.Load()))
}

func (ap *atomicPayload) ConsumeEvent(ev event.EvtLocalAddressesUpdated) error {
	data, err := ev.SignedPeerRecord.Marshal()
	if err == nil {
		(*atomic.Value)(ap).Store(data)
	}

	return err
}

func (ap *atomicPayload) Load() []byte {
	return (*atomic.Value)(ap).Load().([]byte)
}
