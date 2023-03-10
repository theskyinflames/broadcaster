// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package api_test

import (
	"sync"
	"theskyinflames/core-tech/listener/internal/infra/api"
)

// Ensure, that BasicAuthMock does implement api.BasicAuth.
// If this is not the case, regenerate this file with moq.
var _ api.BasicAuth = &BasicAuthMock{}

// BasicAuthMock is a mock implementation of api.BasicAuth.
//
//	func TestSomethingThatUsesBasicAuth(t *testing.T) {
//
//		// make and configure a mocked api.BasicAuth
//		mockedBasicAuth := &BasicAuthMock{
//			AuthFunc: func(user string, pass string) error {
//				panic("mock out the Auth method")
//			},
//		}
//
//		// use mockedBasicAuth in code that requires api.BasicAuth
//		// and then make assertions.
//
//	}
type BasicAuthMock struct {
	// AuthFunc mocks the Auth method.
	AuthFunc func(user string, pass string) error

	// calls tracks calls to the methods.
	calls struct {
		// Auth holds details about calls to the Auth method.
		Auth []struct {
			// User is the user argument value.
			User string
			// Pass is the pass argument value.
			Pass string
		}
	}
	lockAuth sync.RWMutex
}

// Auth calls AuthFunc.
func (mock *BasicAuthMock) Auth(user string, pass string) error {
	callInfo := struct {
		User string
		Pass string
	}{
		User: user,
		Pass: pass,
	}
	mock.lockAuth.Lock()
	mock.calls.Auth = append(mock.calls.Auth, callInfo)
	mock.lockAuth.Unlock()
	if mock.AuthFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.AuthFunc(user, pass)
}

// AuthCalls gets all the calls that were made to Auth.
// Check the length with:
//
//	len(mockedBasicAuth.AuthCalls())
func (mock *BasicAuthMock) AuthCalls() []struct {
	User string
	Pass string
} {
	var calls []struct {
		User string
		Pass string
	}
	mock.lockAuth.RLock()
	calls = mock.calls.Auth
	mock.lockAuth.RUnlock()
	return calls
}
