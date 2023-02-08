package domain

import (
	"errors"
	"net"
	"sync"

	"github.com/google/uuid"
	"github.com/theskyinflames/cqrs-eda/pkg/ddd"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

// Subscriber is self-described
type Subscriber struct {
	ab ddd.AggregateBasic
	net.Conn
}

// NewSubscriber is a constructor
func NewSubscriber(ID uuid.UUID, conn net.Conn) Subscriber {
	return Subscriber{
		ab:   ddd.NewAggregateBasic(ID),
		Conn: conn,
	}
}

// ID returns the underlying ddd.AggregateBasic ID
func (s Subscriber) ID() uuid.UUID {
	return s.ab.ID()
}

// Subscribers is the set of subscribers
type Subscribers struct {
	ab          ddd.AggregateBasic
	mux         *sync.RWMutex
	subscribers map[uuid.UUID]Subscriber
}

// NewSubscribers is a constructor
func NewSubscribers(ID uuid.UUID) Subscribers {
	return Subscribers{
		ab:          ddd.NewAggregateBasic(ID),
		mux:         &sync.RWMutex{},
		subscribers: make(map[uuid.UUID]Subscriber),
	}
}

// Events returns subscribers events
func (s Subscribers) Events() []events.Event {
	return s.ab.Events()
}

// Add adds a new subscriber
func (s *Subscribers) Add(subscriber Subscriber) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.subscribers[subscriber.ID()] = subscriber

	s.ab.RecordEvent(NewSubscriberAddedEvent(subscriber.ID()))
}

// Remove removes a subscriber
func (s *Subscribers) Remove(conn net.Conn) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, subscriber := range s.subscribers {
		if subscriber.Conn.RemoteAddr().String() == conn.RemoteAddr().String() {
			delete(s.subscribers, subscriber.ID())
			s.ab.RecordEvent(NewSubscriberRemovedEvent(subscriber.ID()))
		}
	}
}

// StreamFunction is a stream function
type StreamFunction func(Subscriber)

// Stream applies a StreamFunction function to subscribers
func (s Subscribers) Stream(sl StreamFunction) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, subscriber := range s.subscribers {
		sl(subscriber)
	}
}

// Len returns the quantity of subscribers
func (s Subscribers) Len() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return len(s.subscribers)
}

// ErrSubscriberNotFound is self-described
var ErrSubscriberNotFound = errors.New("subscriber not found")

// Subscriber returns a subscriber by its UUID
func (s *Subscribers) Subscriber(ID uuid.UUID) (Subscriber, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	subscriber, ok := s.subscribers[ID]
	if !ok {
		return Subscriber{}, ErrSubscriberNotFound
	}
	return subscriber, nil
}
