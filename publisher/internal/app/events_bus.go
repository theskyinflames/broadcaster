package app

import (
	"context"
	"errors"
	"fmt"
	"io"

	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/gobwas/ws/wsutil"
	"github.com/theskyinflames/cqrs-eda/pkg/bus"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
	"github.com/theskyinflames/cqrs-eda/pkg/events"
)

// BuildEventsBus returns a generic events bus
func BuildEventsBus(ctx context.Context, commandBus cqrs.Bus, subscribers Subscribers, l cqrs.Logger) bus.Bus {
	eventsBus := bus.New()
	eventsBus.Register(domain.SubscriberAddedEventName, busHandler(subscriberAddedEventHandler(ctx, commandBus, subscribers, l))) // Add subscriber
	eventsBus.Register(domain.SubscriberRemovedEventName, busHandler(eventHandler()))                                             // Remove subscriber
	eventsBus.Register(MessageBroadcastedEventName, busHandler(eventHandler()))                                                   // Message broadcasted
	return eventsBus
}

func busHandler(evh events.Handler) bus.Handler {
	return bus.Handler(func(_ context.Context, d bus.Dispatchable) (interface{}, error) {
		ev, ok := d.(events.Event)
		if !ok {
			return nil, errors.New("is not an event")
		}
		evh(ev)
		return nil, nil
	})
}

func eventHandler() events.Handler {
	return events.Handler(func(ev events.Event) {
		fmt.Printf("received event: %s from aggregate ID: %s\n", ev.Name(), ev.AggregateID().String())
	})
}

func subscriberAddedEventHandler(ctx context.Context, commandBus cqrs.Bus, subscribers Subscribers, l cqrs.Logger) events.Handler {
	return events.Handler(func(ev events.Event) {
		subscriber, err := subscribers.Subscriber(ev.AggregateID())
		if err != nil {
			l.Printf("starting subscriber %s: %d\n", ev.AggregateID().String(), err.Error())
			return
		}
		go func(ctx context.Context, subscriber domain.Subscriber, commandBus cqrs.Bus, l cqrs.Logger) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if _, _, err := wsutil.ReadClientData(subscriber.Conn); err != nil {
						if err == io.EOF {
							// the subscriber has closed its connection
							// the subscriber is removed from subscribers set
							if _, err = commandBus.Dispatch(ctx, RemoveSubscriberCmd{Subscriber: subscriber}); err != nil {
								l.Printf("removing new subscriber: %s\n", err.Error())
							}
							return
						}
					}
				}
			}
		}(ctx, subscriber, commandBus, l)
	})
}
