package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"theskyinflames/core-tech/publisher/internal/app"
	"theskyinflames/core-tech/publisher/internal/domain"
	"theskyinflames/core-tech/publisher/internal/infra/redis"

	api "theskyinflames/core-tech/publisher/internal/infra/http"
	httpx "theskyinflames/core-tech/publisher/internal/infra/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"github.com/theskyinflames/cqrs-eda/pkg/bus"
)

const srvPort = ":81"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	r := chi.NewRouter()

	l := log.New(os.Stdout, "listener: ", os.O_APPEND)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:81"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{"GET"},
	})
	r.Use(cors.Handler)
	r.Use(middleware.Logger)

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	subscribers := domain.NewSubscribers(uuid.New())

	cb := bus.New()
	eb := app.BuildEventsBus(ctx, cb, &subscribers, l)
	cb = app.BuildCommandBus(cb, eb, httpx.SendMessage, &subscribers, l)

	redisConfig, err := redis.NewConfig()
	if err != nil {
		panic(err)
	}
	queue, err := redis.NewQueue(redisConfig)
	if err != nil {
		panic(err)
	}
	go func() { queue.Subscribe(ctx, redisConfig.Topic, cb, l) }()

	HTTPHandler := api.AddSubscriber(ctx, cb, l)

	// The WS HTTP handler is exposed wrapped into the authentication middleware
	r.Get("/websocket", HTTPHandler)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)
	go func() {
		s := <-sig
		cancel()
		l.Printf("signal %q received; listener is shutting down", s.String())
	}()

	l.Printf("serving at port %s\n", srvPort)
	if err := http.ListenAndServe(srvPort, r); err != nil {
		panic(fmt.Sprintf("something went wrong trying to start the server: %s", err.Error()))
	}
}
