package redis

import (
	"context"
	"fmt"
	"log"

	"theskyinflames/core-tech/publisher/internal/app"

	redis "github.com/go-redis/redis/v7"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
)

// ReadErr is a queue error
type ReadErr struct{}

// NewReadErr is a constructor
func NewReadErr(err error) error {
	return fmt.Errorf("reading message: %w", err)
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

// Subscribe subscribes to a topic and wait for messages to be published
func (q Queue) Subscribe(ctx context.Context, topic string, commandBus cqrs.Bus, l *log.Logger) {
	pubsub := q.client.Subscribe(topic)
	defer pubsub.Close()

	// Wait for incoming messages
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				l.Println(NewReadErr(err).Error())
			}
			_, _ = commandBus.Dispatch(ctx, app.BroadcastMessageCmd{Msg: msg.Payload})
		}
	}
}
