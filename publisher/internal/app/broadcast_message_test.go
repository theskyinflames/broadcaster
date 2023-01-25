package app_test

import (
	"context"
	"testing"

	"theskyinflames/core-tech/publisher/internal/app"
	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/stretchr/testify/require"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
)

func TestBroadcastMessage(t *testing.T) {
	tests := []struct {
		name            string
		cmd             cqrs.Command
		broadcaster     app.Broadcaster
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
			name:        `Given message to be broadcasted, when it's called, then it's broadcasted`,
			cmd:         app.BroadcastMessageCmd{},
			broadcaster: func(subscriber domain.Subscriber, msg string) error { return nil },
			subscribers: &SubscribersMock{},
		},
	}

	for _, tt := range tests {
		ch := app.NewBroadcastMessage(tt.broadcaster, tt.subscribers)
		ev, err := ch.Handle(context.Background(), tt.cmd)
		require.Equal(t, tt.expectedErrFunc == nil, err == nil)
		if err != nil {
			tt.expectedErrFunc(t, err)
			continue
		}

		require.Len(t, tt.subscribers.StreamCalls(), 1)
		require.Len(t, ev, 1)
		require.Equal(t, ev[0].Name(), app.MessageBroadcastedEventName)
	}
}
