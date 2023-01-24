package app_test

import (
	"context"
	"errors"
	"testing"

	"theskyinflames/core-tech/listener/internal/app"

	"github.com/stretchr/testify/require"
)

func TestReceiveMessage(t *testing.T) {
	t.Run(`Given a message handler that returns an error on handling a message,
			when it's called with a message,
			then an error is returned`, func(t *testing.T) {
		randomErr := errors.New("")
		mh := &MessageHandlerMock{
			HandleFunc: func(_ context.Context, _ string) error {
				return randomErr
			},
		}

		rm := app.NewReceiveMessage(mh)
		require.ErrorIs(t, randomErr, rm.ReceiveMessage(context.Background(), ""))
	})

	t.Run(`Given a message handler,
	when it's called with a message,
	then no error is returned and the message is passed to the handler`, func(t *testing.T) {
		mh := &MessageHandlerMock{
			HandleFunc: func(_ context.Context, _ string) error {
				return nil
			},
		}

		rm := app.NewReceiveMessage(mh)
		msg := "a_message"
		require.NoError(t, rm.ReceiveMessage(context.Background(), msg))
		require.Len(t, mh.HandleCalls(), 1)
		require.Equal(t, mh.HandleCalls()[0].Msg, msg)
	})
}
