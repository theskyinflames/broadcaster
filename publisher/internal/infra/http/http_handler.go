package api

import (
	"context"
	"log"
	"net/http"

	"theskyinflames/core-tech/publisher/internal/app"
	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/gobwas/ws"
	"github.com/google/uuid"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
)

//go:generate moq -stub -out zmock_api_test.go -pkg api_test . WSHandler

// https://www.blitter.se/utils/basic-authentication-header-generator/
// websocat -H="Authorization: Basic dXNlcmRkOnB3ZA=="  ws://localhost:8080/websocket

// AddSubscriber is the websocket entry-point
func AddSubscriber(ctx context.Context, commandBus cqrs.Bus, l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			l.Printf("get ws connection: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		subscriber := domain.NewSubscriber(uuid.New(), conn)

		// Adding the new subscriber
		if _, err = commandBus.Dispatch(ctx, app.AddSubscriberCmd{
			Subscriber: subscriber,
		}); err != nil {
			l.Panicf("adding new subscriber: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
