package domain_test

import (
	"net"
	"testing"

	"theskyinflames/core-tech/publisher/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSubscribers(t *testing.T) {
	t.Run(`Given a subscribers set`, func(t *testing.T) {
		subscribers := domain.NewSubscribers(uuid.New())

		addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:80")
		require.NoError(t, err)
		mockConn := &mockConn{}
		mockConn.On("RemoteAddr").Return(addr)
		subscriber := domain.NewSubscriber(uuid.New(), mockConn)

		t.Run(`when a new subscriber is added,
				then it's added to the set`, func(t *testing.T) {
			subscribers.Add(subscriber)
			require.Equal(t, subscribers.Len(), 1)
		})

		t.Run(`when it's streamed with a function,
				then this function is applied to subscribers`, func(t *testing.T) {
			var streamed domain.Subscriber
			sf := domain.StreamFunction(func(s domain.Subscriber) { streamed = s })
			subscribers.Stream(sf)
			require.Equal(t, subscriber.ID(), streamed.ID())
		})

		t.Run(`when a subscriber is removed,
				then it does not belongs to the set anymore`, func(t *testing.T) {
			subscribers.Remove(subscriber.Conn)
			require.Equal(t, subscribers.Len(), 0)
		})
	})
}
