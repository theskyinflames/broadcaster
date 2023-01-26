package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"theskyinflames/core-tech/listener/internal/app"
	"theskyinflames/core-tech/listener/internal/infra/api"
	"theskyinflames/core-tech/listener/internal/infra/basicauth"
	"theskyinflames/core-tech/listener/internal/infra/redis"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gobwas/ws/wsutil"

	"github.com/rs/cors"
)

const srvPort = ":80"

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

	// The messageHandler is the application service in charge of manage the received messages
	mh, err := messageHandler()
	if err != nil {
		panic(fmt.Sprintf("connect to messages queue: %s", err.Error()))
	}

	// Here the basic-auth mechanism is set
	authConfig, err := basicauth.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("building auth config: %s", err.Error()))
	}
	basicauth := api.NewBasicAuthFromHTTPRequest(
		basicauth.NewBasicAuth(
			authConfig.Credentials,
		),
	)
	authnMiddleware := api.AuthnMiddleware(basicauth)

	// The conversation handler is in charge of handle the WS conversation fluxes
	wsConversationHandler := api.NewWSConversationHandlerBasic(mh, wsutil.ReadClientData, l)

	// This is the HTTP handler
	wsHTTPHandler := api.MessagesListener(ctx, wsConversationHandler, l)

	// The WS HTTP handler is exposed wrapped into the authentication middleware
	r.Get("/websocket", authnMiddleware(wsHTTPHandler))

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

func messageHandler() (app.MessageHandler, error) {
	c, err := redis.NewConfig()
	if err != nil {
		return nil, err
	}
	return redis.NewQueue(c)
}
