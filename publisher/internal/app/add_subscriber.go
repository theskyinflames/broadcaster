package app

import (
	"context"

	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

// AddSubscriberCmd is a command
type AddSubscriberCmd struct {
	Subscriber domain.Subscriber
}

// AddSubscriberName is self-described
var AddSubscriberName = "add.subscriber"

// Name implements the Command interface
func (cmd AddSubscriberCmd) Name() string {
	return AddSubscriberName
}

// AddSubscriber is a command handler
type AddSubscriber struct {
	subscribers Subscribers
}

// NewAddSubscriber is a constructor
func NewAddSubscriber(subscribers Subscribers) AddSubscriber {
	return AddSubscriber{
		subscribers: subscribers,
	}
}

// Handle implements CommandHandler interface
func (ch AddSubscriber) Handle(ctx context.Context, cmd cqrs.Command) ([]events.Event, error) {
	co, ok := cmd.(AddSubscriberCmd)
	if !ok {
		return nil, NewInvalidCommandError(AddSubscriberName, cmd.Name())
	}

	ch.subscribers.Add(co.Subscriber)

	return ch.subscribers.Events(), nil
}
