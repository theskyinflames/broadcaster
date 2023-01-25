package api

import (
	"context"
	"errors"
	"io"
	"log"
	"net"

	"theskyinflames/core-tech/listener/internal/app"

	"github.com/gobwas/ws"
)

// DataReader read the received messages
type DataReader func(rw io.ReadWriter) ([]byte, ws.OpCode, error)

// WSConversationHandlerBasic manages the ws conversation, from its beginning until its end.
type WSConversationHandlerBasic struct {
	mh app.MessageHandler
	dr DataReader
	l  *log.Logger
}

// NewWSConversationHandlerBasic is a constructor
func NewWSConversationHandlerBasic(mh app.MessageHandler, dr DataReader, l *log.Logger) WSConversationHandlerBasic {
	return WSConversationHandlerBasic{
		mh: mh,
		dr: dr,
		l:  l,
	}
}

// Handle implements the WSConversationHandler interface
func (w WSConversationHandlerBasic) Handle(ctx context.Context, conn net.Conn) error {
	w.l.Printf("broadcaster has open a new ws connection\n")
	for {
		select {
		case <-ctx.Done():
			w.l.Printf("listener finished, closing open ws conversation\n")
			return nil
		default:
			connClosed, err := w.ReceiveMessage(ctx, conn)
			if err != nil {
				return err
			}
			if connClosed {
				return nil
			}
		}
	}
}

// ReceiveMessage receives messages from ws connection and pass them to the messages handler
func (w WSConversationHandlerBasic) ReceiveMessage(ctx context.Context, conn net.Conn) (connClosed bool, err error) {
	b, op, err := w.dr(conn)
	if err != nil {
		if errors.Is(err, io.EOF) {
			w.l.Printf("broadcaster has closed ws connection\n")
			return true, nil
		}
		logErr(w.l, "reading broadcaster message from ws connection", err)
		return false, err
	}
	msg := string(b)

	if op.IsData() {
		if err := w.mh.Handle(ctx, msg); err != nil {
			logErr(w.l, "handle message", err)
			return false, err
		}
	}
	return false, nil
}
