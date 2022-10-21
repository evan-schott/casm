package stream_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"capnproto.org/go/capnp/v3"
	capnp_stream "capnproto.org/go/capnp/v3/std/capnp/stream"
	testing_api "github.com/wetware/casm/internal/api/testing"
	"github.com/wetware/casm/pkg/util/stream"
)

func TestState(t *testing.T) {
	t.Parallel()
	t.Helper()

	t.Run("Succeed", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		server := &streamer{}
		client := testing_api.Streamer_ServerToClient(server)
		defer client.Release()

		s := stream.New(client.Recv)
		s.Call(ctx, nil)
	})

	t.Run("CallAndWait", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		server := sleepStreamer(time.Millisecond)
		client := testing_api.Streamer_ServerToClient(server)
		defer client.Release()

		s := stream.New(client.Recv)

		// stream 10 calls; each blocks for 1ms
		for i := 0; i < 10; i++ {
			s.Call(ctx, nil)
		}

		assert.NoError(t, s.Wait(), "should finish gracefully")
	})

	t.Run("AbortWait", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		server := sleepStreamer(time.Second)
		client := testing_api.Streamer_ServerToClient(server)
		defer client.Release()

		s := stream.New(client.Recv)

		s.Call(ctx, nil)

		cherr := make(chan error, 1)
		go func() {
			cherr <- s.Wait()
		}()

		cancel()

		assert.Eventually(t, func() bool {
			return !s.Open()
		}, time.Millisecond*100, time.Millisecond*10,
			"call to Wait() should signal stream close")

		select {
		case <-time.After(time.Millisecond * 500):
			t.Error("failed to abort after 500ms")
		case err := <-cherr:
			require.ErrorIs(t, err, context.Canceled)
		}
	})

	t.Run("ContextExpired", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		server := &streamer{}
		client := testing_api.Streamer_ServerToClient(server)
		defer client.Release()

		s := stream.New(client.Recv)

		// make one successful call so that the receive-loop is
		// started.
		s.Call(ctx, nil)

		err := s.Wait()
		assert.ErrorIs(t, err, context.Canceled, "error: %v", err)
	})

	t.Run("HandlerError", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		server := &streamer{error: errors.New("test")}
		client := testing_api.Streamer_ServerToClient(server)
		defer client.Release()

		s := stream.New(client.Recv)

		// Make a call so that the stream's receive-loop receives the
		// server error, which should cause the stream to abort.
		s.Call(ctx, nil)

		assert.Error(t, s.Wait(),
			"Wait() should return error from server")
	})
}

func TestStream(t *testing.T) {
	t.Parallel()

	/*
		Test tracking of an actual stream.
	*/

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := counter(0)
	client := testing_api.Streamer_ServerToClient(&server)
	defer client.Release()

	s := stream.New(client.Recv)

	for i := 0; i < 100; i++ {
		require.True(t, s.Open(), "stream should be open")
		s.Call(ctx, nil)
	}

	assert.NoError(t, s.Wait(), "should succeed")
	assert.Equal(t, 100, int(server), "should process 100 calls")
}

func TestStream_NoCall(t *testing.T) {
	t.Parallel()

	s := stream.New(nop)

	assert.NoError(t, s.Wait(), "should succeed")
}

func BenchmarkStream(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// We don't want to benchmark capnp; just the stream manager.
	// The focus is primarily on avoiding garbage from the underlying
	// linked-list.
	s := stream.New(nop)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s.Call(ctx, nil)
	}

	if err := s.Wait(); err != nil {
		panic(err)
	}
}

var nopFuture = capnp_stream.StreamResult_Future{
	Future: capnp.ErrorAnswer(capnp.Method{
		InterfaceID:   0xef96789c0d60cd00,
		MethodID:      0,
		InterfaceName: "testing.capnp:Echoer",
		MethodName:    "echo",
	}, nil).Future(),
}

func nop(context.Context, func(testing_api.Streamer_recv_Params) error) (capnp_stream.StreamResult_Future, capnp.ReleaseFunc) {
	return nopFuture, func() {}
}

type streamer struct{ error }

func (s *streamer) Recv(ctx context.Context, err testing_api.Streamer_recv) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return s.error
}

type counter int

func (ctr *counter) Recv(context.Context, testing_api.Streamer_recv) error {
	*ctr++
	return nil
}

type sleepStreamer time.Duration

func (s sleepStreamer) Recv(ctx context.Context, _ testing_api.Streamer_recv) error {
	select {
	case <-time.After(time.Duration(s)):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
