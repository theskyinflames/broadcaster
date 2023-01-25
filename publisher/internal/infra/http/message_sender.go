package api

import (
	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// SendMessage sends a message to a subscriber
func SendMessage(subscriber domain.Subscriber, msg string) error {
	return wsutil.WriteServerMessage(subscriber.Conn, ws.OpText, []byte(msg))
}
