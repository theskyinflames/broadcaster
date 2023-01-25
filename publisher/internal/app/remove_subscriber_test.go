package app_test

import (
	"context"
	"net"
	"testing"

	"theskyinflames/core-tech/publisher/internal/app"
	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
)

func TestRemoveSubscriber(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:80")
	require.NoError(t, err)
	connMock := new(mockConn)
	connMock.On("LocalAddr").Return(addr)

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
			name: `Given message to be broadcasted, when it's called, then it's broadcasted`,
			cmd: app.RemoveSubscriberCmd{
				Subscriber: domain.NewSubscriber(uuid.New(), connMock),
			},
			subscribers: &SubscribersMock{},
		},
	}

	for _, tt := range tests {
		ch := app.NewRemoveSubscriber(tt.subscribers)
		_, err := ch.Handle(context.Background(), tt.cmd)
		require.Equal(t, tt.expectedErrFunc == nil, err == nil)
		if err != nil {
			tt.expectedErrFunc(t, err)
			continue
		}

		require.Len(t, tt.subscribers.RemoveCalls(), 1)
		require.Equal(t,
			tt.cmd.(app.RemoveSubscriberCmd).Subscriber.Conn.LocalAddr().String(),
			tt.subscribers.RemoveCalls()[0].Conn.LocalAddr().String())
	}
}
