package app

import (
	"context"

	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

// RemoveSubscriberCmd is a command
type RemoveSubscriberCmd struct {
	Subscriber domain.Subscriber
}

// RemoveSubscriberName is self-described
var RemoveSubscriberName = "Remove.subscriber"

// Name implements the Command interface
func (cmd RemoveSubscriberCmd) Name() string {
	return RemoveSubscriberName
}

// RemoveSubscriber is a command handler
type RemoveSubscriber struct {
	subscribers Subscribers
}

// NewRemoveSubscriber is a constructor
func NewRemoveSubscriber(subscribers Subscribers) RemoveSubscriber {
	return RemoveSubscriber{
		subscribers: subscribers,
	}
}

// Handle implements CommandHandler interface
func (ch RemoveSubscriber) Handle(ctx context.Context, cmd cqrs.Command) ([]events.Event, error) {
	co, ok := cmd.(RemoveSubscriberCmd)
	if !ok {
		return nil, NewInvalidCommandError(RemoveSubscriberName, cmd.Name())
	}

	ch.subscribers.Remove(co.Subscriber)

	return ch.subscribers.Events(), nil
}
