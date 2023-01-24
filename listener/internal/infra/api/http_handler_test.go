package api_test

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"theskyinflames/core-tech/listener/internal/infra/api"

	hijackable "github.com/getlantern/httptest"
	"github.com/stretchr/testify/require"
)

type writerMock struct {
	mux   *sync.Mutex
	calls int
}

func (w *writerMock) Write(_ []byte) (n int, err error) {
	w.mux.Lock()
	defer w.mux.Unlock()

	w.calls++
	return
}

func (w *writerMock) callsValue() int {
	w.mux.Lock()
	defer w.mux.Unlock()

	return w.calls
}

func TestMessagesListener(t *testing.T) {
	t.Run(`Given an not hijackable request,
	when it's called,
	then the error is logged and an HTTP error code is returned `, func(t *testing.T) {
		ctx := context.Background()
		writer := &writerMock{mux: &sync.Mutex{}}
		hf := api.MessagesListener(
			ctx,
			&WSConversationHandlerMock{
				HandleFunc: func(ctx context.Context, conn net.Conn) error {
					return errors.New("")
				},
			},
			log.New(writer, "listener: ", os.O_APPEND),
		)

		srv := httptest.NewServer(http.HandlerFunc(hf))
		host, _, _ := net.SplitHostPort(srv.Listener.Addr().String())
		defer srv.Close()

		req, _ := http.NewRequest(http.MethodGet, "/websocket", nil)
		req.Header.Add("Authorization", "Basic dXNlcjpwd2Q=")
		req.Host = host

		rr := hijackable.NewRecorder([]byte{})
		hf.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code())
		waitForErrorLogged(ctx, t, writer)
	})

	t.Run(`Given an websocket conversation handler that returns an error,
	when it's called,
	then the conversation finishes and the error is logged`, func(t *testing.T) {
		writer := &writerMock{mux: &sync.Mutex{}}
		ctx := context.Background()
		hf := api.MessagesListener(
			ctx,
			&WSConversationHandlerMock{
				HandleFunc: func(ctx context.Context, conn net.Conn) error {
					return errors.New("")
				},
			},
			log.New(writer, "listener: ", os.O_APPEND),
		)

		srv := httptest.NewServer(http.HandlerFunc(hf))
		host, _, _ := net.SplitHostPort(srv.Listener.Addr().String())
		defer srv.Close()

		req, _ := http.NewRequest(http.MethodGet, "/websocket", nil)
		req.Header.Add("Authorization", "Basic dXNlcjpwd2Q=")
		req.Header.Add("Connection", "Upgrade")
		req.Header.Add("Upgrade", "websocket")
		req.Header.Add("Sec-WebSocket-Version", "13")
		req.Header.Add("Sec-WebSocket-Key", "q4xkcO32u266gldTuKaSOw==")
		req.Host = host

		rr := hijackable.NewRecorder([]byte{})
		hf.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code())
		waitForErrorLogged(ctx, t, writer)
	})
}

func waitForErrorLogged(ctx context.Context, t *testing.T, w *writerMock) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

endloop:
	for {
		select {
		case <-ctx.Done():
			t.Fail()
		default:
			if w.callsValue() == 1 {
				break endloop
			}
		}
	}
	require.Equal(t, 1, w.calls)
}
