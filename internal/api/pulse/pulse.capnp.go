// Code generated by capnpc-go. DO NOT EDIT.

package pulse

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Heartbeat capnp.Struct

// Heartbeat_TypeID is the unique identifier for the type Heartbeat.
const Heartbeat_TypeID = 0x83bca0e4d70e0e82

func NewHeartbeat(s *capnp.Segment) (Heartbeat, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Heartbeat(st), err
}

func NewRootHeartbeat(s *capnp.Segment) (Heartbeat, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Heartbeat(st), err
}

func ReadRootHeartbeat(msg *capnp.Message) (Heartbeat, error) {
	root, err := msg.Root()
	return Heartbeat(root.Struct()), err
}

func (s Heartbeat) String() string {
	str, _ := text.Marshal(0x83bca0e4d70e0e82, capnp.Struct(s))
	return str
}

func (s Heartbeat) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Heartbeat) DecodeFromPtr(p capnp.Ptr) Heartbeat {
	return Heartbeat(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Heartbeat) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Heartbeat) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Heartbeat) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Heartbeat) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Heartbeat) Ttl() uint32 {
	return capnp.Struct(s).Uint32(0)
}

func (s Heartbeat) SetTtl(v uint32) {
	capnp.Struct(s).SetUint32(0, v)
}

func (s Heartbeat) Id() uint32 {
	return capnp.Struct(s).Uint32(4)
}

func (s Heartbeat) SetId(v uint32) {
	capnp.Struct(s).SetUint32(4, v)
}

func (s Heartbeat) Hostname() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Heartbeat) HasHostname() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Heartbeat) HostnameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Heartbeat) SetHostname(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Heartbeat) Meta() (Heartbeat_Field_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Heartbeat_Field_List(p.List()), err
}

func (s Heartbeat) HasMeta() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Heartbeat) SetMeta(v Heartbeat_Field_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewMeta sets the meta field to a newly
// allocated Heartbeat_Field_List, preferring placement in s's segment.
func (s Heartbeat) NewMeta(n int32) (Heartbeat_Field_List, error) {
	l, err := NewHeartbeat_Field_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Heartbeat_Field_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}

// Heartbeat_List is a list of Heartbeat.
type Heartbeat_List = capnp.StructList[Heartbeat]

// NewHeartbeat creates a new list of Heartbeat.
func NewHeartbeat_List(s *capnp.Segment, sz int32) (Heartbeat_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[Heartbeat](l), err
}

// Heartbeat_Future is a wrapper for a Heartbeat promised by a client call.
type Heartbeat_Future struct{ *capnp.Future }

func (p Heartbeat_Future) Struct() (Heartbeat, error) {
	s, err := p.Future.Struct()
	return Heartbeat(s), err
}

type Heartbeat_Field capnp.Struct

// Heartbeat_Field_TypeID is the unique identifier for the type Heartbeat_Field.
const Heartbeat_Field_TypeID = 0xd79cdf84dc4551cb

func NewHeartbeat_Field(s *capnp.Segment) (Heartbeat_Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Heartbeat_Field(st), err
}

func NewRootHeartbeat_Field(s *capnp.Segment) (Heartbeat_Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Heartbeat_Field(st), err
}

func ReadRootHeartbeat_Field(msg *capnp.Message) (Heartbeat_Field, error) {
	root, err := msg.Root()
	return Heartbeat_Field(root.Struct()), err
}

func (s Heartbeat_Field) String() string {
	str, _ := text.Marshal(0xd79cdf84dc4551cb, capnp.Struct(s))
	return str
}

func (s Heartbeat_Field) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Heartbeat_Field) DecodeFromPtr(p capnp.Ptr) Heartbeat_Field {
	return Heartbeat_Field(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Heartbeat_Field) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Heartbeat_Field) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Heartbeat_Field) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Heartbeat_Field) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Heartbeat_Field) Key() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Heartbeat_Field) HasKey() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Heartbeat_Field) KeyBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Heartbeat_Field) SetKey(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Heartbeat_Field) Value() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Heartbeat_Field) HasValue() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Heartbeat_Field) ValueBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Heartbeat_Field) SetValue(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

// Heartbeat_Field_List is a list of Heartbeat_Field.
type Heartbeat_Field_List = capnp.StructList[Heartbeat_Field]

// NewHeartbeat_Field creates a new list of Heartbeat_Field.
func NewHeartbeat_Field_List(s *capnp.Segment, sz int32) (Heartbeat_Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Heartbeat_Field](l), err
}

// Heartbeat_Field_Future is a wrapper for a Heartbeat_Field promised by a client call.
type Heartbeat_Field_Future struct{ *capnp.Future }

func (p Heartbeat_Field_Future) Struct() (Heartbeat_Field, error) {
	s, err := p.Future.Struct()
	return Heartbeat_Field(s), err
}

const schema_c2974e3dc137fcee = "x\xdaL\x8f1K#Q\x14\x85\xcfyo\xb2\x93\"" +
	"\x9b\xcd#a\xb3l\x13\x10\x0b\x15\x124\x08\x82 \xd8" +
	"\x18\xc4Br{\x0bG\xf3 \x83\x93\x18\xcdD\xb1U" +
	"\x0b\xff\x81\x8d\x85\xbfA\xb0\xb3\xd1\xd6\x1f\x90F\x10D" +
	"\x10\x1b!\x85\x85 \x8cL`\x92t\xe7\x1d\xde\xbd\xf7" +
	"\xfbr\xb7\xab\xce\xc2\xef\xa2\x82\x92\x7f\xa9_\xd1i6" +
	"\xdb\x7f\xbd\xbe;\x83d\xc8\xe8\xe3{\xe9~e\xf3\xf2" +
	"\x01)\xe5\x02fpc\xbe\x8a@\x9e|\x03\xa3GY" +
	"{:\x7f\xbe\xea\xc3\xfc\xe5xn\xf83\xff\xce\x97\xfc" +
	"'\xe34\xe01\xcaQ\xa7\x17tme\xd7c\xa7\xdd" +
	"Y^\xb7\xdea)\xdc\xb1^(\x0e'\xf7\xb0Z\xaa" +
	"\xf96hHN;\x80C\xc0xS\x80liJS" +
	"\xd1\x90\x05\xc6\xa5\xfd\x0f\xc8\xb6\xa6\x04\x8aT\x05*\xc0" +
	"\xf8\x1b\x8045%T4\x9a\x05j\xc0\x1c\xcc\x01\x12" +
	"h\xca\x85\xa2\x1b\x86\x01\xd3PL\x83\xdao$1j" +
	"\xeew\xc3\xb6\xd7\xb2\x00\x98\x81b\x06\xfc\xd3\xb2\xa1\xc7" +
	",X\xd7dn\x0c\x08\xc6\xe5HF%2C\x97J" +
	"\xcdwm\xd0\xa8\x93\x92\x1e\xe1\xcf\xc6\xf8\xd3\x9a2?" +
	"\x81_\xae\x022\xa3)\x8b\x8a\xee\x9e=I\xce\x96\x8e" +
	"\xbc\xa0g\x93\xd7O\x00\x00\x00\xff\xffl\xda^\x7f"

func init() {
	schemas.Register(schema_c2974e3dc137fcee,
		0x83bca0e4d70e0e82,
		0xd79cdf84dc4551cb)
}
