package app

import (
	"net"

	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/google/uuid"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

//go:generate moq -stub -out zmock_app_subscribers_test.go -pkg app_test . Subscribers

// Subscribers gathers the set of subscribers
type Subscribers interface {
	Events() []events.Event

	Add(subscriber domain.Subscriber)
	Remove(conn net.Conn)
	Stream(sl domain.StreamFunction)
	Subscriber(uuid.UUID) (domain.Subscriber, error)
	Len() int
}
