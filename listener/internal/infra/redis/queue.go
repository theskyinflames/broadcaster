package redis

import (
	"context"
	"fmt"
	"net/http"

	redis "github.com/go-redis/redis/v7"
)

// BasicAuth provides basic-authentication based on http request
type BasicAuth interface {
	func(r *http.Request) error
}

// PublishErr is a queue error
type PublishErr struct{}

// NewPublishErr is a constructor
func NewPublishErr(err error) error {
	return fmt.Errorf("publishing message: %w", err)
}

// Queue is in charge of manage the messages queue
type Queue struct {
	topic  string
	client *redis.Client
}

// NewQueue is a constructor
func NewQueue(c Config) (Queue, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,     // p.e "localhost:6379",
		Password: c.Password, // "" means no password set
		DB:       c.Db,       // 0 to use default DB
	})
	if err := client.Ping().Err(); err != nil {
		return Queue{}, fmt.Errorf("redis not available: %w", err)
	}
	return Queue{
		topic:  c.Topic,
		client: client,
	}, nil
}

// Handle is in charge of handle a received message trough websocket api
func (q Queue) Handle(ctx context.Context, msg string) error {
	fmt.Printf("sending message: %s\n", msg)
	if err := q.client.Publish(q.topic, msg).Err(); err != nil {
		return NewPublishErr(err)
	}
	return nil
}
