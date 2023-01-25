package app

import (
	"github.com/google/uuid"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

// MessageBroadcastedEventName is self-described
const MessageBroadcastedEventName = "message.broadcasted"

// MessageBroadcastedEvent is an event
type MessageBroadcastedEvent struct {
	events.EventBasic
}

// NewMessageBroadcastedEvent is a constructor
func NewMessageBroadcastedEvent(msg string) MessageBroadcastedEvent {
	return MessageBroadcastedEvent{
		EventBasic: events.NewEventBasic(uuid.New(), MessageBroadcastedEventName, msg),
	}
}
