// Code generated by capnpc-go. DO NOT EDIT.

package pex

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Gossip capnp.Struct

// Gossip_TypeID is the unique identifier for the type Gossip.
const Gossip_TypeID = 0xfcef4b0e93332397

func NewGossip(s *capnp.Segment) (Gossip, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Gossip(st), err
}

func NewRootGossip(s *capnp.Segment) (Gossip, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Gossip(st), err
}

func ReadRootGossip(msg *capnp.Message) (Gossip, error) {
	root, err := msg.Root()
	return Gossip(root.Struct()), err
}

func (s Gossip) String() string {
	str, _ := text.Marshal(0xfcef4b0e93332397, capnp.Struct(s))
	return str
}

func (s Gossip) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Gossip) DecodeFromPtr(p capnp.Ptr) Gossip {
	return Gossip(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Gossip) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Gossip) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Gossip) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Gossip) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Gossip) Hop() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Gossip) SetHop(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Gossip) Envelope() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s Gossip) HasEnvelope() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Gossip) SetEnvelope(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

// Gossip_List is a list of Gossip.
type Gossip_List = capnp.StructList[Gossip]

// NewGossip creates a new list of Gossip.
func NewGossip_List(s *capnp.Segment, sz int32) (Gossip_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Gossip](l), err
}

// Gossip_Future is a wrapper for a Gossip promised by a client call.
type Gossip_Future struct{ *capnp.Future }

func (f Gossip_Future) Struct() (Gossip, error) {
	p, err := f.Future.Ptr()
	return Gossip(p.Struct()), err
}

const schema_bbd81c151780f030 = "x\xda\x12Ht`1\xe4\xdd\xcf\xc8\xc0\x14(\xc2\xca" +
	"\xf6\x7f\xba\xb2\xf1d>\xef\xf7\x7f\x18\x02\xb9\x18\x19\xff" +
	"\x1b|h\x10\x17\x95\xb9\xb1\x9b\x81\x95\x91\x9d\x81A\xf0" +
	"h\x97\xe0Y\x10}\xb2\x9cA\xf7\x7fAj\x85^r" +
	"bA\x1ec\x81\x95{~qqf\x01\x03C\x00#" +
	"c \x073\x0b\x03\x03\x0b#\x03\x83\xa0\xa6\x12\x03C" +
	"\xa0\x0a3c\xa0\x01\x13##\xa3\x08#HL\xd7\x8b" +
	"\x81!P\x87\x991\xd0\x82\x89\x91=#\xbf\x80\x91\x93" +
	"\x81\x89\x91\x93\x81\xf1\x7fj^YjN~A*\x03" +
	"\x03\x03#/\x03\x13#/\x03# \x00\x00\xff\xfft" +
	"w j"

func init() {
	schemas.Register(schema_bbd81c151780f030,
		0xfcef4b0e93332397)
}
