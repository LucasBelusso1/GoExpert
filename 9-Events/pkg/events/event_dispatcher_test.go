package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dispacher *EventDispatcher
var event TestEvent
var event2 TestEvent
var handler TestEventHandler
var handler2 TestEventHandler
var handler3 TestEventHandler

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}

	dispacher = suite.eventDispatcher
	event = suite.event
	event2 = suite.event2
	handler = suite.handler
	handler2 = suite.handler2
	handler3 = suite.handler3
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := dispacher.Register(event.GetName(), &handler)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event.GetName()]))
	suite.Equal(&handler, dispacher.handlers[event.GetName()][0])

	err = dispacher.Register(event.GetName(), &handler2)
	suite.Nil(err)
	suite.Equal(2, len(dispacher.handlers[event.GetName()]))
	suite.Equal(&handler2, dispacher.handlers[event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := dispacher.Register(event.GetName(), &handler)
	suite.Nil(err)
	err = dispacher.Register(event.GetName(), &handler)
	suite.Error(err)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 1
	err := dispacher.Register(event.GetName(), &handler)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event.GetName()]))

	err = dispacher.Register(event.GetName(), &handler2)
	suite.Nil(err)
	suite.Equal(2, len(dispacher.handlers[event.GetName()]))

	// Event 2
	err = dispacher.Register(event2.GetName(), &handler3)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event2.GetName()]))

	dispacher.Clear()
	suite.Nil(err)
	suite.Empty(dispacher.handlers)
	suite.Equal(0, len(dispacher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := dispacher.Register(event.GetName(), &handler)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event.GetName()]))

	hasEvent := dispacher.Has(event.GetName(), &handler)
	suite.True(hasEvent)

	err = dispacher.Register(event.GetName(), &handler2)
	suite.Nil(err)
	suite.Equal(2, len(dispacher.handlers[event.GetName()]))

	hasEvent = dispacher.Has(event.GetName(), &handler2)
	suite.True(hasEvent)

	hasEvent = dispacher.Has(event.GetName(), &handler3)
	suite.False(hasEvent)

	hasEvent = dispacher.Has(event2.GetName(), &handler3)
	suite.False(hasEvent)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := dispacher.Register(event.GetName(), &handler)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event.GetName()]))

	err = dispacher.Register(event.GetName(), &handler2)
	suite.Nil(err)
	suite.Equal(2, len(dispacher.handlers[event.GetName()]))

	err = dispacher.Register(event2.GetName(), &handler3)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event2.GetName()]))

	dispacher.Remove(event.GetName(), &handler)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event.GetName()]))
	suite.Equal(&handler2, dispacher.handlers[event.GetName()][0])

	dispacher.Remove(event.GetName(), &handler)
	suite.Error(ErrHandlerNotFound)

	dispacher.Remove(event.GetName(), &handler2)
	suite.Nil(err)
	suite.Equal(0, len(dispacher.handlers[event.GetName()]))

	dispacher.Remove(event2.GetName(), &handler3)
	suite.Nil(err)
	suite.Equal(0, len(dispacher.handlers[event2.GetName()]))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &event)

	dispacher.Register(event.GetName(), eh)
	dispacher.Register(event.GetName(), eh2)
	dispacher.Dispatch(&event)
	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
