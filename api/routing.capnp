using Go = import "/go.capnp";

@0xc2974e3dc137fcee;

$Go.package("routing");
$Go.import("github.com/wetware/casm/internal/api/routing");

using Milliseconds = UInt32;


struct Heartbeat {
    # Heartbeat messages are used to implement an unstructured p2p
    # clustering service.  Hosts periodically emit heartbeats on a
    # pubsub topic (the "namespace") and construct a routing table
    # based on heartbeats received by other peers.
    #
    # Additional metadata can piggyback off of heartbeat messages,
    # allowing indexed operations on the routing table.

    ttl      @0 :Milliseconds;
    # Time-to-live, in milliseconds. The originator is considered
    # failed if a subsequent heartbeat is not received within ttl.
    
    server @1 :UInt64;
    # An opaque identifier that uniquely distinguishes an instance
    # of a host. This identifier is randomly generated each time a
    # host boots.

    host     @2 :Text;
    # The hostname of the underlying host, as reported by the OS.
    # Users MUST NOT assume hostnames to be unique or non-empty.

    meta     @3 :List(Text);
    # A set of optional, arbitrary metadata fields.  Fields are
    # encoded as key-value pairs separated by the '=' rune.  Fields
    # are parsed into keys and values by splitting the string on the
    # first occurrenc of the '=' separator.  Subsequent occurrences
    # are treated as part of the value.
}


interface View {
    # A View is a read-only snapshot of a particular host's routing
    # table. Views are not updated, and should therefore be queried
    # and discarded promptly.

    lookup  @0 (selector :Selector, constraints :List(Constraint)) -> (result :MaybeRecord);
    iter    @1 (handler :Handler, selector :Selector, constraints :List(Constraint)) -> ();
    reverse @2 () -> (view :View);
    
    interface Handler {
        recv @0 (record :Record) -> stream;
    }

    struct Selector {
        union {
            all        @0 :Void;
            match      @1 :Index;
            from       @2 :Index;
        }
    }

    struct Constraint {
        union {
            limit      @0 :UInt64;
            to         @1 :Index;
        }
    }

    struct Index {
        prefix    @0 :Bool;
        union {
            id    @1 :PeerID;
            host  @2 :Text;
            meta  @3 :Text;        # key=value
        }
    }

    struct Record {
        peer      @0 :PeerID;
        seq       @1 :UInt64;
        heartbeat @2 :Heartbeat;
    }

    struct MaybeRecord {
        union {
            nothing @0 :Void;
            just    @1 :Record;
        }
    }

    using PeerID = Text;
}
