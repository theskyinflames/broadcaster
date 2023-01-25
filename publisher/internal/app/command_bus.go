package app

import (
	"github.com/theskyinflames/cqrs-eda/pkg/bus"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
	"github.com/theskyinflames/cqrs-eda/pkg/helpers"
)

// BuildCommandBus returns the command bus
func BuildCommandBus(bus bus.Bus, eventsBus bus.Bus, msgSender Broadcaster, subscribers Subscribers, l cqrs.Logger) bus.Bus {
	chMw := cqrs.CommandHandlerMultiMiddleware(
		cqrs.ChEventMw(eventsBus),
		cqrs.ChErrMw(l),
	)

	broadcastMessage := chMw(NewBroadcastMessage(msgSender, subscribers))
	addSubscriber := chMw(NewAddSubscriber(subscribers))
	removeSubscriber := chMw(NewRemoveSubscriber(subscribers))

	bus.Register(BroadcastMessageName, helpers.BusChHandler(broadcastMessage))
	bus.Register(AddSubscriberName, helpers.BusChHandler(addSubscriber))
	bus.Register(RemoveSubscriberName, helpers.BusChHandler(removeSubscriber))
	return bus
}
