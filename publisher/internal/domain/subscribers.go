package domain

import (
	"errors"
	"net"
	"sync"

	"github.com/google/uuid"
	"github.com/theskyinflames/cqrs-eda/pkg/ddd"
)

// Subscriber is self-described
type Subscriber struct {
	ddd.AggregateBasic
	net.Conn
}

// NewSubscriber is a constructor
func NewSubscriber(ID uuid.UUID, conn net.Conn) Subscriber {
	return Subscriber{
		AggregateBasic: ddd.NewAggregateBasic(ID),
		Conn:           conn,
	}
}

// Subscribers is the set of subscribers
type Subscribers struct {
	ddd.AggregateBasic
	mux         *sync.RWMutex
	subscribers map[uuid.UUID]Subscriber
}

// NewSubscribers is a constructor
func NewSubscribers(ID uuid.UUID) Subscribers {
	return Subscribers{
		AggregateBasic: ddd.NewAggregateBasic(ID),
		mux:            &sync.RWMutex{},
		subscribers:    make(map[uuid.UUID]Subscriber),
	}
}

// Add adds a new subscriber
func (s *Subscribers) Add(subscriber Subscriber) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.subscribers[subscriber.ID()] = subscriber

	s.RecordEvent(NewSubscriberAddedEvent(subscriber.ID()))
}

// Remove removes a subscriber
func (s *Subscribers) Remove(conn net.Conn) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, subscriber := range s.subscribers {
		if subscriber.Conn.RemoteAddr().String() == conn.RemoteAddr().String() {
			delete(s.subscribers, subscriber.ID())
			s.RecordEvent(NewSubscriberRemovedEvent(subscriber.ID()))
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
