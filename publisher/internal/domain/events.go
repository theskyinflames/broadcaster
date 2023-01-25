package domain

import (
	"github.com/google/uuid"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

// SubscriberAddedEventName is self-described
const SubscriberAddedEventName = "subscriber.added"

// SubscriberAddedEvent is an event
type SubscriberAddedEvent struct {
	events.EventBasic
}

// NewSubscriberAddedEvent is a constructor
func NewSubscriberAddedEvent(ID uuid.UUID) SubscriberAddedEvent {
	return SubscriberAddedEvent{
		EventBasic: events.NewEventBasic(ID, SubscriberAddedEventName, ID.String()),
	}
}

// SubscriberRemovedEventName is self-described
const SubscriberRemovedEventName = "subscriber.removed"

// SubscriberRemovedEvent is an event
type SubscriberRemovedEvent struct {
	events.EventBasic
}

// NewSubscriberRemovedEvent is a constructor
func NewSubscriberRemovedEvent(ID uuid.UUID) SubscriberRemovedEvent {
	return SubscriberRemovedEvent{
		EventBasic: events.NewEventBasic(ID, SubscriberRemovedEventName, ID.String()),
	}
}
