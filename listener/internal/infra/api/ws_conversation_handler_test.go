package api_test

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"theskyinflames/core-tech/listener/internal/infra/api"

	"github.com/gobwas/ws"
	"github.com/stretchr/testify/require"
)

func TestReceiveMessage(t *testing.T) {
	l := log.New(&writerMock{mux: &sync.Mutex{}}, "", os.O_APPEND)
	t.Run(`Given a data reader that returns an error,
			when it's called,
			then an error is returned`, func(t *testing.T) {
		mh := &MessageHandlerMock{
			HandleFunc: func(_ context.Context, _ string) error {
				return nil
			},
		}
		randomErr := errors.New("")
		dr := api.DataReader(func(rw io.ReadWriter) ([]byte, ws.OpCode, error) {
			return nil, '\\', randomErr
		})

		wsch := api.NewWSConversationHandlerBasic(mh, dr, l)
		closed, err := wsch.ReceiveMessage(context.Background(), nil)
		require.ErrorIs(t, err, randomErr)
		require.False(t, closed)
	})

	t.Run(`Given a data reader that returns an io.EOF error,
			when it's called,
			then the connection is closed`, func(t *testing.T) {
		mh := &MessageHandlerMock{
			HandleFunc: func(_ context.Context, _ string) error {
				return nil
			},
		}
		dr := api.DataReader(func(rw io.ReadWriter) ([]byte, ws.OpCode, error) {
			return nil, '\\', io.EOF
		})

		wsch := api.NewWSConversationHandlerBasic(mh, dr, l)
		closed, err := wsch.ReceiveMessage(context.Background(), nil)
		require.NoError(t, err)
		require.True(t, closed)
	})

	t.Run(`Given a data message handler that returns an error,
			when it's called,
			then an error is returned`, func(t *testing.T) {
		randomErr := errors.New("")
		mh := &MessageHandlerMock{
			HandleFunc: func(_ context.Context, _ string) error {
				return randomErr
			},
		}
		msg := []byte("msg")
		dr := api.DataReader(func(rw io.ReadWriter) ([]byte, ws.OpCode, error) {
			return msg, ws.OpText, nil
		})

		wsch := api.NewWSConversationHandlerBasic(mh, dr, l)
		closed, err := wsch.ReceiveMessage(context.Background(), nil)
		require.ErrorIs(t, err, randomErr)
		require.False(t, closed)
		require.Len(t, mh.HandleCalls(), 1)
		require.Equal(t, string(msg), mh.HandleCalls()[0].Msg)
	})

	t.Run(`Given a received message,
			when it's called,
			then no error is returned`, func(t *testing.T) {
		mh := &MessageHandlerMock{
			HandleFunc: func(_ context.Context, _ string) error {
				return nil
			},
		}
		msg := []byte("msg")
		dr := api.DataReader(func(rw io.ReadWriter) ([]byte, ws.OpCode, error) {
			return msg, ws.OpText, nil
		})

		wsch := api.NewWSConversationHandlerBasic(mh, dr, l)
		closed, err := wsch.ReceiveMessage(context.Background(), nil)
		require.NoError(t, err)
		require.False(t, closed)
		require.Len(t, mh.HandleCalls(), 1)
		require.Equal(t, string(msg), mh.HandleCalls()[0].Msg)
	})
}
