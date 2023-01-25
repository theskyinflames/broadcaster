package app_test

import (
	"context"
	"testing"

	"theskyinflames/core-tech/publisher/internal/app"
	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
)

func newInvalidCommand() cqrs.Command {
	return &CommandMock{
		NameFunc: func() string {
			return "invalid_command"
		},
	}
}

func TestAddSubscriber(t *testing.T) {
	tests := []struct {
		name            string
		cmd             cqrs.Command
		subscribers     *SubscribersMock
		expectedErrFunc func(*testing.T, error)
	}{
		{
			name: `Given an invalid command, when it's called, then an error is returned`,
			cmd:  newInvalidCommand(),
			expectedErrFunc: func(t *testing.T, err error) {
				require.ErrorAs(t, err, &app.InvalidCommandError{})
			},
		},
		{
			name:        `Given command, when it's called, then the subscriber is added`,
			cmd:         app.AddSubscriberCmd{Subscriber: domain.NewSubscriber(uuid.New(), nil)},
			subscribers: &SubscribersMock{},
		},
	}

	for _, tt := range tests {
		ch := app.NewAddSubscriber(tt.subscribers)
		_, err := ch.Handle(context.Background(), tt.cmd)
		require.Equal(t, tt.expectedErrFunc == nil, err == nil)
		if err != nil {
			tt.expectedErrFunc(t, err)
			continue
		}

		require.Len(t, tt.subscribers.AddCalls(), 1)
		require.Equal(t, tt.cmd.(app.AddSubscriberCmd).Subscriber.ID(), tt.subscribers.AddCalls()[0].Subscriber.ID())
	}
}
