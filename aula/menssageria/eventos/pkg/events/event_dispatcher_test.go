package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (te *TestEvent) GetName() string {
	return te.Name
}

func (te *TestEvent) GetPayload() interface{} {
	return te.Payload
}

func (te *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	// Implementação do tratamento do evento para fins de teste
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event1          TestEvent
	event2          TestEvent
	handler1        *TestEventHandler
	handler2        *TestEventHandler
	handler3        *TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler1 = &TestEventHandler{ID: 1}
	suite.handler2 = &TestEventHandler{ID: 2}
	suite.handler3 = &TestEventHandler{ID: 3}
	suite.event1 = TestEvent{Name: "Event1", Payload: "Payload1"}
	suite.event2 = TestEvent{Name: "Event2", Payload: "Payload2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	err = suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	assert.Equal(suite.T(), suite.handler1, suite.eventDispatcher.handlers[suite.event1.GetName()][0])
	assert.Equal(suite.T(), suite.handler2, suite.eventDispatcher.handlers[suite.event1.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	err = suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler1)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Clear() {
	// event 1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	err = suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	err = suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event1.GetName(), suite.handler1))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event1.GetName(), suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event1.GetName(), suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{} // mock
	// eh.On -> configura o mock para esperar a chamada do método Handle com o evento suite.event1
	eh.On("Handle", &suite.event1)
	suite.eventDispatcher.Register(suite.event1.GetName(), eh)

	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event1)
	suite.eventDispatcher.Register(suite.event1.GetName(), eh2)

	err := suite.eventDispatcher.Dispatch(&suite.event1)
	suite.Nil(err)

	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	err = suite.eventDispatcher.Register(suite.event1.GetName(), suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	// Remover handler1
	err = suite.eventDispatcher.Remove(suite.event1.GetName(), suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))
	suite.Equal(suite.handler2, suite.eventDispatcher.handlers[suite.event1.GetName()][0])

	err = suite.eventDispatcher.Remove(suite.event1.GetName(), suite.handler2)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// Remover handler3
	err = suite.eventDispatcher.Remove(suite.event2.GetName(), suite.handler3)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
