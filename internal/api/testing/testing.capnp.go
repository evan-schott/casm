// Code generated by capnpc-go. DO NOT EDIT.

package testing

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	stream "capnproto.org/go/capnp/v3/std/capnp/stream"
	context "context"
	fmt "fmt"
)

type Echoer capnp.Client

// Echoer_TypeID is the unique identifier for the type Echoer.
const Echoer_TypeID = 0xef96789c0d60cd00

func (c Echoer) Echo(ctx context.Context, params func(Echoer_echo_Params) error) (Echoer_echo_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xef96789c0d60cd00,
			MethodID:      0,
			InterfaceName: "testing.capnp:Echoer",
			MethodName:    "echo",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Echoer_echo_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Echoer_echo_Results_Future{Future: ans.Future()}, release
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Echoer) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Echoer) AddRef() Echoer {
	return Echoer(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Echoer) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Echoer) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Echoer) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Echoer) DecodeFromPtr(p capnp.Ptr) Echoer {
	return Echoer(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Echoer) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Echoer) IsSame(other Echoer) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Echoer) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Echoer) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
} // A Echoer_Server is a Echoer with a local implementation.
type Echoer_Server interface {
	Echo(context.Context, Echoer_echo) error
}

// Echoer_NewServer creates a new Server from an implementation of Echoer_Server.
func Echoer_NewServer(s Echoer_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Echoer_Methods(nil, s), s, c)
}

// Echoer_ServerToClient creates a new Client from an implementation of Echoer_Server.
// The caller is responsible for calling Release on the returned Client.
func Echoer_ServerToClient(s Echoer_Server) Echoer {
	return Echoer(capnp.NewClient(Echoer_NewServer(s)))
}

// Echoer_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Echoer_Methods(methods []server.Method, s Echoer_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xef96789c0d60cd00,
			MethodID:      0,
			InterfaceName: "testing.capnp:Echoer",
			MethodName:    "echo",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Echo(ctx, Echoer_echo{call})
		},
	})

	return methods
}

// Echoer_echo holds the state for a server call to Echoer.echo.
// See server.Call for documentation.
type Echoer_echo struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Echoer_echo) Args() Echoer_echo_Params {
	return Echoer_echo_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Echoer_echo) AllocResults() (Echoer_echo_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echoer_echo_Results(r), err
}

// Echoer_List is a list of Echoer.
type Echoer_List = capnp.CapList[Echoer]

// NewEchoer creates a new list of Echoer.
func NewEchoer_List(s *capnp.Segment, sz int32) (Echoer_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Echoer](l), err
}

type Echoer_echo_Params capnp.Struct

// Echoer_echo_Params_TypeID is the unique identifier for the type Echoer_echo_Params.
const Echoer_echo_Params_TypeID = 0x97878b56a59dadff

func NewEchoer_echo_Params(s *capnp.Segment) (Echoer_echo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echoer_echo_Params(st), err
}

func NewRootEchoer_echo_Params(s *capnp.Segment) (Echoer_echo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echoer_echo_Params(st), err
}

func ReadRootEchoer_echo_Params(msg *capnp.Message) (Echoer_echo_Params, error) {
	root, err := msg.Root()
	return Echoer_echo_Params(root.Struct()), err
}

func (s Echoer_echo_Params) String() string {
	str, _ := text.Marshal(0x97878b56a59dadff, capnp.Struct(s))
	return str
}

func (s Echoer_echo_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Echoer_echo_Params) DecodeFromPtr(p capnp.Ptr) Echoer_echo_Params {
	return Echoer_echo_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Echoer_echo_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Echoer_echo_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Echoer_echo_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Echoer_echo_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Echoer_echo_Params) Payload() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Echoer_echo_Params) HasPayload() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Echoer_echo_Params) PayloadBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Echoer_echo_Params) SetPayload(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// Echoer_echo_Params_List is a list of Echoer_echo_Params.
type Echoer_echo_Params_List = capnp.StructList[Echoer_echo_Params]

// NewEchoer_echo_Params creates a new list of Echoer_echo_Params.
func NewEchoer_echo_Params_List(s *capnp.Segment, sz int32) (Echoer_echo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Echoer_echo_Params](l), err
}

// Echoer_echo_Params_Future is a wrapper for a Echoer_echo_Params promised by a client call.
type Echoer_echo_Params_Future struct{ *capnp.Future }

func (f Echoer_echo_Params_Future) Struct() (Echoer_echo_Params, error) {
	p, err := f.Future.Ptr()
	return Echoer_echo_Params(p.Struct()), err
}

type Echoer_echo_Results capnp.Struct

// Echoer_echo_Results_TypeID is the unique identifier for the type Echoer_echo_Results.
const Echoer_echo_Results_TypeID = 0xa46d5fd1187fd6ee

func NewEchoer_echo_Results(s *capnp.Segment) (Echoer_echo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echoer_echo_Results(st), err
}

func NewRootEchoer_echo_Results(s *capnp.Segment) (Echoer_echo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echoer_echo_Results(st), err
}

func ReadRootEchoer_echo_Results(msg *capnp.Message) (Echoer_echo_Results, error) {
	root, err := msg.Root()
	return Echoer_echo_Results(root.Struct()), err
}

func (s Echoer_echo_Results) String() string {
	str, _ := text.Marshal(0xa46d5fd1187fd6ee, capnp.Struct(s))
	return str
}

func (s Echoer_echo_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Echoer_echo_Results) DecodeFromPtr(p capnp.Ptr) Echoer_echo_Results {
	return Echoer_echo_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Echoer_echo_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Echoer_echo_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Echoer_echo_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Echoer_echo_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Echoer_echo_Results) Result() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Echoer_echo_Results) HasResult() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Echoer_echo_Results) ResultBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Echoer_echo_Results) SetResult(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// Echoer_echo_Results_List is a list of Echoer_echo_Results.
type Echoer_echo_Results_List = capnp.StructList[Echoer_echo_Results]

// NewEchoer_echo_Results creates a new list of Echoer_echo_Results.
func NewEchoer_echo_Results_List(s *capnp.Segment, sz int32) (Echoer_echo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Echoer_echo_Results](l), err
}

// Echoer_echo_Results_Future is a wrapper for a Echoer_echo_Results promised by a client call.
type Echoer_echo_Results_Future struct{ *capnp.Future }

func (f Echoer_echo_Results_Future) Struct() (Echoer_echo_Results, error) {
	p, err := f.Future.Ptr()
	return Echoer_echo_Results(p.Struct()), err
}

type Streamer capnp.Client

// Streamer_TypeID is the unique identifier for the type Streamer.
const Streamer_TypeID = 0x87f59ceaa25ddecc

func (c Streamer) Recv(ctx context.Context, params func(Streamer_recv_Params) error) (stream.StreamResult_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x87f59ceaa25ddecc,
			MethodID:      0,
			InterfaceName: "testing.capnp:Streamer",
			MethodName:    "recv",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Streamer_recv_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return stream.StreamResult_Future{Future: ans.Future()}, release
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Streamer) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Streamer) AddRef() Streamer {
	return Streamer(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Streamer) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Streamer) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Streamer) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Streamer) DecodeFromPtr(p capnp.Ptr) Streamer {
	return Streamer(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Streamer) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Streamer) IsSame(other Streamer) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Streamer) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Streamer) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
} // A Streamer_Server is a Streamer with a local implementation.
type Streamer_Server interface {
	Recv(context.Context, Streamer_recv) error
}

// Streamer_NewServer creates a new Server from an implementation of Streamer_Server.
func Streamer_NewServer(s Streamer_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Streamer_Methods(nil, s), s, c)
}

// Streamer_ServerToClient creates a new Client from an implementation of Streamer_Server.
// The caller is responsible for calling Release on the returned Client.
func Streamer_ServerToClient(s Streamer_Server) Streamer {
	return Streamer(capnp.NewClient(Streamer_NewServer(s)))
}

// Streamer_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Streamer_Methods(methods []server.Method, s Streamer_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x87f59ceaa25ddecc,
			MethodID:      0,
			InterfaceName: "testing.capnp:Streamer",
			MethodName:    "recv",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Recv(ctx, Streamer_recv{call})
		},
	})

	return methods
}

// Streamer_recv holds the state for a server call to Streamer.recv.
// See server.Call for documentation.
type Streamer_recv struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Streamer_recv) Args() Streamer_recv_Params {
	return Streamer_recv_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Streamer_recv) AllocResults() (stream.StreamResult, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return stream.StreamResult(r), err
}

// Streamer_List is a list of Streamer.
type Streamer_List = capnp.CapList[Streamer]

// NewStreamer creates a new list of Streamer.
func NewStreamer_List(s *capnp.Segment, sz int32) (Streamer_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Streamer](l), err
}

type Streamer_recv_Params capnp.Struct

// Streamer_recv_Params_TypeID is the unique identifier for the type Streamer_recv_Params.
const Streamer_recv_Params_TypeID = 0xf01f1821166adede

func NewStreamer_recv_Params(s *capnp.Segment) (Streamer_recv_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Streamer_recv_Params(st), err
}

func NewRootStreamer_recv_Params(s *capnp.Segment) (Streamer_recv_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Streamer_recv_Params(st), err
}

func ReadRootStreamer_recv_Params(msg *capnp.Message) (Streamer_recv_Params, error) {
	root, err := msg.Root()
	return Streamer_recv_Params(root.Struct()), err
}

func (s Streamer_recv_Params) String() string {
	str, _ := text.Marshal(0xf01f1821166adede, capnp.Struct(s))
	return str
}

func (s Streamer_recv_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Streamer_recv_Params) DecodeFromPtr(p capnp.Ptr) Streamer_recv_Params {
	return Streamer_recv_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Streamer_recv_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Streamer_recv_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Streamer_recv_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Streamer_recv_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Streamer_recv_Params_List is a list of Streamer_recv_Params.
type Streamer_recv_Params_List = capnp.StructList[Streamer_recv_Params]

// NewStreamer_recv_Params creates a new list of Streamer_recv_Params.
func NewStreamer_recv_Params_List(s *capnp.Segment, sz int32) (Streamer_recv_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Streamer_recv_Params](l), err
}

// Streamer_recv_Params_Future is a wrapper for a Streamer_recv_Params promised by a client call.
type Streamer_recv_Params_Future struct{ *capnp.Future }

func (f Streamer_recv_Params_Future) Struct() (Streamer_recv_Params, error) {
	p, err := f.Future.Ptr()
	return Streamer_recv_Params(p.Struct()), err
}

const schema_86c7b3eb31eb86de = "x\xda\x12x\xe2\xc0b\xc8[\xcf\xc2\xc0\x14h\xc2\xca" +
	"\xf6\xff\xcc\xbd\xd8E\xaf\xe6|mg\x10\xe4c\xfe\x7f" +
	"\xaf\xed\xb5\xe1\xeb\xcd\xc7\xdb\x18\x18\x18\x85U\x99v\x09" +
	"\xeb2\xb130\x08k2\xb9\x0b\x07\x82X\xff\xff\xaf" +
	"\x9d\xbb4\xac\xbb}:\x83\xa0(#\x03\x03+#;" +
	"\x03\x83\xb1%\x13\x17#\x03\xa3\xb0#\x93=\x03\xe3\xff" +
	"w\xd7\xea%.\xc6\xe7.AV\x90\xc8$\x04R\x90" +
	"\x09R\xf0\xefl\x02\xef\x9c\x8ai\xef\xd1-\xebdZ" +
	"%<\x11lY/\x93\xbb\xf0V\xb0e\xf7\xeee\x89" +
	")J\xc8\x7f`\x10\x14gd``\x01\x195\x97I" +
	"\x8a\x91!\xe6\x7fIjqIf^\xba\x1ecrb" +
	"A^\x81Up\x89}QjbnjQ\x00#c" +
	" \x0b3+\x92^\xc6\xbc\x8d\x07\xca\x8dg\xc5\xcf\x14" +
	"\x14\xd4b`\x12de\xe7/JM.s`\x0c`" +
	"d\x84\x9b\xc2\x0c1\xc559#?\xb5H/59" +
	"#_% \xb1(1\xb7\x98\x81!\x90\x85\x99\x85\x81" +
	"\x81\x85\x91\x81A\x90\xd7\x89\x81!\x90\x83\x991P\x84" +
	"\x89\xb1\xbe \xb12'?1\x85\x91\x87\x81\x89\x91\x87" +
	"\x01\xafQA\xa9\xc5\xa59%\x8c\xc5\xc8FY!\x8c" +
	"\xb2/\x82\xc8\xa3\x9b\xc4\x083\x89\x1fd\x14\xc2c\xb0" +
	"\x18`\x84\x854\xdcc \xdb\xb0z,\xb8\x04\x12:" +
	"z \xaf\x83\xbd\xc6\x9c[\x0c\x08\x00\x00\xff\xff`$" +
	"\x94\xb0"

func init() {
	schemas.Register(schema_86c7b3eb31eb86de,
		0x87f59ceaa25ddecc,
		0x97878b56a59dadff,
		0xa46d5fd1187fd6ee,
		0xef96789c0d60cd00,
		0xf01f1821166adede)
}
