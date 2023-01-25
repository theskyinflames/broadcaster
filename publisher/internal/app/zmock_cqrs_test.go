// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package app_test

import (
	"sync"
	"github.com/theskyinflames/cqrs-eda/pkg/cqrs"
)

// Ensure, that CommandMock does implement cqrs.Command.
// If this is not the case, regenerate this file with moq.
var _ cqrs.Command = &CommandMock{}

// CommandMock is a mock implementation of cqrs.Command.
//
//	func TestSomethingThatUsesCommand(t *testing.T) {
//
//		// make and configure a mocked cqrs.Command
//		mockedCommand := &CommandMock{
//			NameFunc: func() string {
//				panic("mock out the Name method")
//			},
//		}
//
//		// use mockedCommand in code that requires cqrs.Command
//		// and then make assertions.
//
//	}
type CommandMock struct {
	// NameFunc mocks the Name method.
	NameFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Name holds details about calls to the Name method.
		Name []struct {
		}
	}
	lockName sync.RWMutex
}

// Name calls NameFunc.
func (mock *CommandMock) Name() string {
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	if mock.NameFunc == nil {
		var (
			sOut string
		)
		return sOut
	}
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//
//	len(mockedCommand.NameCalls())
func (mock *CommandMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

