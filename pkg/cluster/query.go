package cluster

import (
	"errors"
	"fmt"

	api "github.com/wetware/casm/internal/api/routing"
	"github.com/wetware/casm/pkg/cluster/routing"
)

/*
	Selectors
*/

func All() Selector {
	return func(s SelectorStruct) error {
		s.SetAll()
		return nil
	}
}

func Match(index routing.Index) Selector {
	return func(s SelectorStruct) error {
		return bindIndex(s.NewMatch, index)
	}
}

func From(index routing.Index) Selector {
	return func(s SelectorStruct) error {
		return bindIndex(s.NewFrom, index)
	}
}

/*
	Helpers
*/

func bindIndex(fn func() (api.View_Index, error), index routing.Index) error {
	target, err := fn()
	if err != nil {
		return err
	}

	target.SetPrefix(index.Prefix())

	switch index.String() {
	case "id":
		return bindPeer(target, index)

	case "host":
		return bindHost(target, index)

	case "meta":
		return bindMeta(target, index)
	}

	return fmt.Errorf("invalid index: %s", index)
}

func bindPeer(target api.View_Index, index routing.Index) error {
	switch ix := index.(type) {
	case routing.PeerIndex:
		b, err := ix.PeerBytes()
		if err == nil {
			return target.SetId(string(b)) // TODO:  unsafe.Pointer
		}
		return err

	case interface{ Peer() (string, error) }:
		id, err := ix.Peer()
		if err == nil {
			err = target.SetId(id)
		}
		return err
	}

	return errors.New("not a peer index")
}

func bindHost(target api.View_Index, index routing.Index) error {
	switch ix := index.(type) {
	case routing.HostIndex:
		b, err := ix.HostBytes()
		if err == nil {
			return target.SetHost(string(b)) // TODO:  unsafe.Pointer
		}
		return err

	case interface{ Host() (string, error) }:
		id, err := ix.Host()
		if err == nil {
			err = target.SetHost(id)
		}
		return err
	}

	return errors.New("not a peer index")
}

func bindMeta(target api.View_Index, index routing.Index) error {
	switch ix := index.(type) {
	case routing.MetaIndex:
		b, err := ix.MetaBytes()
		if err == nil {
			err = target.SetMeta(string(b))
		}
		return err

	case interface{ MetaField() (routing.Field, error) }:
		f, err := ix.MetaField()
		if err == nil {
			err = target.SetMeta(f.String())
		}
		return err
	}

	return errors.New("not a metadata index")
}
