package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
)

//go:generate moq -stub -out zmock_api_test.go -pkg api_test . WSConversationHandler

type WSConversationHandler interface {
	Handle(ctx context.Context, conn net.Conn) error
}

// https://www.blitter.se/utils/basic-authentication-header-generator/
// websocat -H="Authorization: Basic dXNlcmRkOnB3ZA=="  ws://localhost:8080/websocket

// MessagesListener is the websocket entry-point
func MessagesListener(ctx context.Context, wsch WSConversationHandler, l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// open connection for ws
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			l.Printf("get ws connection: %s", err.Error())
			return
		}

		// Start ws conversation
		go func(l *log.Logger, conn net.Conn, wsch WSConversationHandler) {
			defer conn.Close()
			if err := wsch.Handle(ctx, conn); err != nil {
				l.Printf("handle ws conversation: %s\n", err.Error())
			}
		}(l, conn, wsch)
	}
}

func logErr(l *log.Logger, when string, err error) {
	l.Println(fmt.Errorf("%s: %w", when, err).Error())
}
