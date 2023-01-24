package app

import (
	"context"
)

//go:generate moq -stub -out zmock_app_message_handler_test.go -pkg app_test . MessageHandler

// MessageHandler handles a received message
type MessageHandler interface {
	Handle(ctx context.Context, msg string) error
}

// ReceiveMessage is a DTO
type ReceiveMessage struct {
	mh MessageHandler
}

// NewReceiveMessage is a constructor
func NewReceiveMessage(mh MessageHandler) ReceiveMessage {
	return ReceiveMessage{mh: mh}
}

// ReceiveMessage implements the use case of receiving a message
func (rm ReceiveMessage) ReceiveMessage(ctx context.Context, msg string) error {
	return rm.mh.Handle(ctx, msg)
}
