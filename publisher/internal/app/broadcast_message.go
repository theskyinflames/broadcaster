package app

import (
	"context"
	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

// Broadcaster sends a message to a subscriber
type Broadcaster func(subscriber domain.Subscriber, msg string) error

// BroadcastMessageCmd is a command
type BroadcastMessageCmd struct {
	Msg string
}

// BroadcastMessageName is self-described
var BroadcastMessageName = "broadcast.message"

// Name implements the Command interface
func (cmd BroadcastMessageCmd) Name() string {
	return BroadcastMessageName
}

// BroadcastMessage is a command handler
type BroadcastMessage struct {
	broadcaster Broadcaster
	subscribers Subscribers
}

// NewBroadcastMessage is a constructor
func NewBroadcastMessage(broadcaster Broadcaster, subscribers Subscribers) BroadcastMessage {
	return BroadcastMessage{
		broadcaster: broadcaster,
		subscribers: subscribers,
	}
}

// Handle implements CommandHandler interface
func (ch BroadcastMessage) Handle(ctx context.Context, cmd cqrs.Command) ([]events.Event, error) {
	co, ok := cmd.(BroadcastMessageCmd)
	if !ok {
		return nil, NewInvalidCommandError(BroadcastMessageName, cmd.Name())
	}

	ch.subscribers.Stream(func(s domain.Subscriber) {
		_ = ch.broadcaster(s, co.Msg)
	})

	return []events.Event{NewMessageBroadcastedEvent(co.Msg)}, nil
}
