// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package db_mock

import (
	apimodels "github.com/keptn/go-utils/pkg/api/models"
	"github.com/keptn/keptn/shipyard-controller/common"
	"github.com/keptn/keptn/shipyard-controller/models"
	"sync"
)

// EventRepoMock is a mock implementation of db.EventRepo.
//
// 	func TestSomethingThatUsesEventRepo(t *testing.T) {
//
// 		// make and configure a mocked db.EventRepo
// 		mockedEventRepo := &EventRepoMock{
// 			DeleteAllFinishedEventsFunc: func(eventScope models.EventScope) error {
// 				panic("mock out the DeleteAllFinishedEvents method")
// 			},
// 			DeleteEventFunc: func(project string, eventID string, status common.EventStatus) error {
// 				panic("mock out the DeleteEvent method")
// 			},
// 			DeleteEventCollectionsFunc: func(project string) error {
// 				panic("mock out the DeleteEventCollections method")
// 			},
// 			GetEventsFunc: func(project string, filter common.EventFilter, status ...common.EventStatus) ([]models.Event, error) {
// 				panic("mock out the GetEvents method")
// 			},
// 			GetEventsWithRetryFunc: func(project string, filter common.EventFilter, status common.EventStatus, nrRetries int) ([]models.Event, error) {
// 				panic("mock out the GetEventsWithRetry method")
// 			},
// 			GetFinishedEventsFunc: func(eventScope models.EventScope) ([]models.Event, error) {
// 				panic("mock out the GetFinishedEvents method")
// 			},
// 			GetRootEventsFunc: func(params models.GetRootEventParams) (*models.GetEventsResult, error) {
// 				panic("mock out the GetRootEvents method")
// 			},
// 			GetStartedEventsForTriggeredIDFunc: func(eventScope models.EventScope) ([]models.Event, error) {
// 				panic("mock out the GetStartedEventsForTriggeredID method")
// 			},
// 			GetTaskSequenceTriggeredEventFunc: func(eventScope models.EventScope, taskSequenceName string) (*models.Event, error) {
// 				panic("mock out the GetTaskSequenceTriggeredEvent method")
// 			},
// 			InsertEventFunc: func(project string, event models.Event, status common.EventStatus) error {
// 				panic("mock out the InsertEvent method")
// 			},
// 		}
//
// 		// use mockedEventRepo in code that requires db.EventRepo
// 		// and then make assertions.
//
// 	}
type EventRepoMock struct {
	// DeleteAllFinishedEventsFunc mocks the DeleteAllFinishedEvents method.
	DeleteAllFinishedEventsFunc func(eventScope models.EventScope) error

	// DeleteEventFunc mocks the DeleteEvent method.
	DeleteEventFunc func(project string, eventID string, status common.EventStatus) error

	// DeleteEventCollectionsFunc mocks the DeleteEventCollections method.
	DeleteEventCollectionsFunc func(project string) error

	// GetEventsFunc mocks the GetEvents method.
	GetEventsFunc func(project string, filter common.EventFilter, status ...common.EventStatus) ([]apimodels.KeptnContextExtendedCE, error)

	// GetEventsWithRetryFunc mocks the GetEventsWithRetry method.
	GetEventsWithRetryFunc func(project string, filter common.EventFilter, status common.EventStatus, nrRetries int) ([]apimodels.KeptnContextExtendedCE, error)

	// GetFinishedEventsFunc mocks the GetFinishedEvents method.
	GetFinishedEventsFunc func(eventScope models.EventScope) ([]apimodels.KeptnContextExtendedCE, error)

	// GetRootEventsFunc mocks the GetRootEvents method.
	GetRootEventsFunc func(params models.GetRootEventParams) (*models.GetEventsResult, error)

	// GetStartedEventsForTriggeredIDFunc mocks the GetStartedEventsForTriggeredID method.
	GetStartedEventsForTriggeredIDFunc func(eventScope models.EventScope) ([]apimodels.KeptnContextExtendedCE, error)

	// GetTaskSequenceTriggeredEventFunc mocks the GetTaskSequenceTriggeredEvent method.
	GetTaskSequenceTriggeredEventFunc func(eventScope models.EventScope, taskSequenceName string) (*apimodels.KeptnContextExtendedCE, error)

	// InsertEventFunc mocks the InsertEvent method.
	InsertEventFunc func(project string, event apimodels.KeptnContextExtendedCE, status common.EventStatus) error

	// calls tracks calls to the methods.
	calls struct {
		// DeleteAllFinishedEvents holds details about calls to the DeleteAllFinishedEvents method.
		DeleteAllFinishedEvents []struct {
			// EventScope is the eventScope argument value.
			EventScope models.EventScope
		}
		// DeleteEvent holds details about calls to the DeleteEvent method.
		DeleteEvent []struct {
			// Project is the project argument value.
			Project string
			// EventID is the eventID argument value.
			EventID string
			// Status is the status argument value.
			Status common.EventStatus
		}
		// DeleteEventCollections holds details about calls to the DeleteEventCollections method.
		DeleteEventCollections []struct {
			// Project is the project argument value.
			Project string
		}
		// GetEvents holds details about calls to the GetEvents method.
		GetEvents []struct {
			// Project is the project argument value.
			Project string
			// Filter is the filter argument value.
			Filter common.EventFilter
			// Status is the status argument value.
			Status []common.EventStatus
		}
		// GetEventsWithRetry holds details about calls to the GetEventsWithRetry method.
		GetEventsWithRetry []struct {
			// Project is the project argument value.
			Project string
			// Filter is the filter argument value.
			Filter common.EventFilter
			// Status is the status argument value.
			Status common.EventStatus
			// NrRetries is the nrRetries argument value.
			NrRetries int
		}
		// GetFinishedEvents holds details about calls to the GetFinishedEvents method.
		GetFinishedEvents []struct {
			// EventScope is the eventScope argument value.
			EventScope models.EventScope
		}
		// GetRootEvents holds details about calls to the GetRootEvents method.
		GetRootEvents []struct {
			// Params is the params argument value.
			Params models.GetRootEventParams
		}
		// GetStartedEventsForTriggeredID holds details about calls to the GetStartedEventsForTriggeredID method.
		GetStartedEventsForTriggeredID []struct {
			// EventScope is the eventScope argument value.
			EventScope models.EventScope
		}
		// GetTaskSequenceTriggeredEvent holds details about calls to the GetTaskSequenceTriggeredEvent method.
		GetTaskSequenceTriggeredEvent []struct {
			// EventScope is the eventScope argument value.
			EventScope models.EventScope
			// TaskSequenceName is the taskSequenceName argument value.
			TaskSequenceName string
		}
		// InsertEvent holds details about calls to the InsertEvent method.
		InsertEvent []struct {
			// Project is the project argument value.
			Project string
			//models.KeptnContextExtendedCE is the event argument value.
			Event apimodels.KeptnContextExtendedCE
			// Status is the status argument value.
			Status common.EventStatus
		}
	}
	lockDeleteAllFinishedEvents        sync.RWMutex
	lockDeleteEvent                    sync.RWMutex
	lockDeleteEventCollections         sync.RWMutex
	lockGetEvents                      sync.RWMutex
	lockGetEventsWithRetry             sync.RWMutex
	lockGetFinishedEvents              sync.RWMutex
	lockGetRootEvents                  sync.RWMutex
	lockGetStartedEventsForTriggeredID sync.RWMutex
	lockGetTaskSequenceTriggeredEvent  sync.RWMutex
	lockInsertEvent                    sync.RWMutex
}

// DeleteAllFinishedEvents calls DeleteAllFinishedEventsFunc.
func (mock *EventRepoMock) DeleteAllFinishedEvents(eventScope models.EventScope) error {
	if mock.DeleteAllFinishedEventsFunc == nil {
		panic("EventRepoMock.DeleteAllFinishedEventsFunc: method is nil but EventRepo.DeleteAllFinishedEvents was just called")
	}
	callInfo := struct {
		EventScope models.EventScope
	}{
		EventScope: eventScope,
	}
	mock.lockDeleteAllFinishedEvents.Lock()
	mock.calls.DeleteAllFinishedEvents = append(mock.calls.DeleteAllFinishedEvents, callInfo)
	mock.lockDeleteAllFinishedEvents.Unlock()
	return mock.DeleteAllFinishedEventsFunc(eventScope)
}

// DeleteAllFinishedEventsCalls gets all the calls that were made to DeleteAllFinishedEvents.
// Check the length with:
//     len(mockedEventRepo.DeleteAllFinishedEventsCalls())
func (mock *EventRepoMock) DeleteAllFinishedEventsCalls() []struct {
	EventScope models.EventScope
} {
	var calls []struct {
		EventScope models.EventScope
	}
	mock.lockDeleteAllFinishedEvents.RLock()
	calls = mock.calls.DeleteAllFinishedEvents
	mock.lockDeleteAllFinishedEvents.RUnlock()
	return calls
}

// DeleteEvent calls DeleteEventFunc.
func (mock *EventRepoMock) DeleteEvent(project string, eventID string, status common.EventStatus) error {
	if mock.DeleteEventFunc == nil {
		panic("EventRepoMock.DeleteEventFunc: method is nil but EventRepo.DeleteEvent was just called")
	}
	callInfo := struct {
		Project string
		EventID string
		Status  common.EventStatus
	}{
		Project: project,
		EventID: eventID,
		Status:  status,
	}
	mock.lockDeleteEvent.Lock()
	mock.calls.DeleteEvent = append(mock.calls.DeleteEvent, callInfo)
	mock.lockDeleteEvent.Unlock()
	return mock.DeleteEventFunc(project, eventID, status)
}

// DeleteEventCalls gets all the calls that were made to DeleteEvent.
// Check the length with:
//     len(mockedEventRepo.DeleteEventCalls())
func (mock *EventRepoMock) DeleteEventCalls() []struct {
	Project string
	EventID string
	Status  common.EventStatus
} {
	var calls []struct {
		Project string
		EventID string
		Status  common.EventStatus
	}
	mock.lockDeleteEvent.RLock()
	calls = mock.calls.DeleteEvent
	mock.lockDeleteEvent.RUnlock()
	return calls
}

// DeleteEventCollections calls DeleteEventCollectionsFunc.
func (mock *EventRepoMock) DeleteEventCollections(project string) error {
	if mock.DeleteEventCollectionsFunc == nil {
		panic("EventRepoMock.DeleteEventCollectionsFunc: method is nil but EventRepo.DeleteEventCollections was just called")
	}
	callInfo := struct {
		Project string
	}{
		Project: project,
	}
	mock.lockDeleteEventCollections.Lock()
	mock.calls.DeleteEventCollections = append(mock.calls.DeleteEventCollections, callInfo)
	mock.lockDeleteEventCollections.Unlock()
	return mock.DeleteEventCollectionsFunc(project)
}

// DeleteEventCollectionsCalls gets all the calls that were made to DeleteEventCollections.
// Check the length with:
//     len(mockedEventRepo.DeleteEventCollectionsCalls())
func (mock *EventRepoMock) DeleteEventCollectionsCalls() []struct {
	Project string
} {
	var calls []struct {
		Project string
	}
	mock.lockDeleteEventCollections.RLock()
	calls = mock.calls.DeleteEventCollections
	mock.lockDeleteEventCollections.RUnlock()
	return calls
}

// GetEvents calls GetEventsFunc.
func (mock *EventRepoMock) GetEvents(project string, filter common.EventFilter, status ...common.EventStatus) ([]apimodels.KeptnContextExtendedCE, error) {
	if mock.GetEventsFunc == nil {
		panic("EventRepoMock.GetEventsFunc: method is nil but EventRepo.GetEvents was just called")
	}
	callInfo := struct {
		Project string
		Filter  common.EventFilter
		Status  []common.EventStatus
	}{
		Project: project,
		Filter:  filter,
		Status:  status,
	}
	mock.lockGetEvents.Lock()
	mock.calls.GetEvents = append(mock.calls.GetEvents, callInfo)
	mock.lockGetEvents.Unlock()
	return mock.GetEventsFunc(project, filter, status...)
}

// GetEventsCalls gets all the calls that were made to GetEvents.
// Check the length with:
//     len(mockedEventRepo.GetEventsCalls())
func (mock *EventRepoMock) GetEventsCalls() []struct {
	Project string
	Filter  common.EventFilter
	Status  []common.EventStatus
} {
	var calls []struct {
		Project string
		Filter  common.EventFilter
		Status  []common.EventStatus
	}
	mock.lockGetEvents.RLock()
	calls = mock.calls.GetEvents
	mock.lockGetEvents.RUnlock()
	return calls
}

// GetEventsWithRetry calls GetEventsWithRetryFunc.
func (mock *EventRepoMock) GetEventsWithRetry(project string, filter common.EventFilter, status common.EventStatus, nrRetries int) ([]apimodels.KeptnContextExtendedCE, error) {
	if mock.GetEventsWithRetryFunc == nil {
		panic("EventRepoMock.GetEventsWithRetryFunc: method is nil but EventRepo.GetEventsWithRetry was just called")
	}
	callInfo := struct {
		Project   string
		Filter    common.EventFilter
		Status    common.EventStatus
		NrRetries int
	}{
		Project:   project,
		Filter:    filter,
		Status:    status,
		NrRetries: nrRetries,
	}
	mock.lockGetEventsWithRetry.Lock()
	mock.calls.GetEventsWithRetry = append(mock.calls.GetEventsWithRetry, callInfo)
	mock.lockGetEventsWithRetry.Unlock()
	return mock.GetEventsWithRetryFunc(project, filter, status, nrRetries)
}

// GetEventsWithRetryCalls gets all the calls that were made to GetEventsWithRetry.
// Check the length with:
//     len(mockedEventRepo.GetEventsWithRetryCalls())
func (mock *EventRepoMock) GetEventsWithRetryCalls() []struct {
	Project   string
	Filter    common.EventFilter
	Status    common.EventStatus
	NrRetries int
} {
	var calls []struct {
		Project   string
		Filter    common.EventFilter
		Status    common.EventStatus
		NrRetries int
	}
	mock.lockGetEventsWithRetry.RLock()
	calls = mock.calls.GetEventsWithRetry
	mock.lockGetEventsWithRetry.RUnlock()
	return calls
}

// GetFinishedEvents calls GetFinishedEventsFunc.
func (mock *EventRepoMock) GetFinishedEvents(eventScope models.EventScope) ([]apimodels.KeptnContextExtendedCE, error) {
	if mock.GetFinishedEventsFunc == nil {
		panic("EventRepoMock.GetFinishedEventsFunc: method is nil but EventRepo.GetFinishedEvents was just called")
	}
	callInfo := struct {
		EventScope models.EventScope
	}{
		EventScope: eventScope,
	}
	mock.lockGetFinishedEvents.Lock()
	mock.calls.GetFinishedEvents = append(mock.calls.GetFinishedEvents, callInfo)
	mock.lockGetFinishedEvents.Unlock()
	return mock.GetFinishedEventsFunc(eventScope)
}

// GetFinishedEventsCalls gets all the calls that were made to GetFinishedEvents.
// Check the length with:
//     len(mockedEventRepo.GetFinishedEventsCalls())
func (mock *EventRepoMock) GetFinishedEventsCalls() []struct {
	EventScope models.EventScope
} {
	var calls []struct {
		EventScope models.EventScope
	}
	mock.lockGetFinishedEvents.RLock()
	calls = mock.calls.GetFinishedEvents
	mock.lockGetFinishedEvents.RUnlock()
	return calls
}

// GetRootEvents calls GetRootEventsFunc.
func (mock *EventRepoMock) GetRootEvents(params models.GetRootEventParams) (*models.GetEventsResult, error) {
	if mock.GetRootEventsFunc == nil {
		panic("EventRepoMock.GetRootEventsFunc: method is nil but EventRepo.GetRootEvents was just called")
	}
	callInfo := struct {
		Params models.GetRootEventParams
	}{
		Params: params,
	}
	mock.lockGetRootEvents.Lock()
	mock.calls.GetRootEvents = append(mock.calls.GetRootEvents, callInfo)
	mock.lockGetRootEvents.Unlock()
	return mock.GetRootEventsFunc(params)
}

// GetRootEventsCalls gets all the calls that were made to GetRootEvents.
// Check the length with:
//     len(mockedEventRepo.GetRootEventsCalls())
func (mock *EventRepoMock) GetRootEventsCalls() []struct {
	Params models.GetRootEventParams
} {
	var calls []struct {
		Params models.GetRootEventParams
	}
	mock.lockGetRootEvents.RLock()
	calls = mock.calls.GetRootEvents
	mock.lockGetRootEvents.RUnlock()
	return calls
}

// GetStartedEventsForTriggeredID calls GetStartedEventsForTriggeredIDFunc.
func (mock *EventRepoMock) GetStartedEventsForTriggeredID(eventScope models.EventScope) ([]apimodels.KeptnContextExtendedCE, error) {
	if mock.GetStartedEventsForTriggeredIDFunc == nil {
		panic("EventRepoMock.GetStartedEventsForTriggeredIDFunc: method is nil but EventRepo.GetStartedEventsForTriggeredID was just called")
	}
	callInfo := struct {
		EventScope models.EventScope
	}{
		EventScope: eventScope,
	}
	mock.lockGetStartedEventsForTriggeredID.Lock()
	mock.calls.GetStartedEventsForTriggeredID = append(mock.calls.GetStartedEventsForTriggeredID, callInfo)
	mock.lockGetStartedEventsForTriggeredID.Unlock()
	return mock.GetStartedEventsForTriggeredIDFunc(eventScope)
}

// GetStartedEventsForTriggeredIDCalls gets all the calls that were made to GetStartedEventsForTriggeredID.
// Check the length with:
//     len(mockedEventRepo.GetStartedEventsForTriggeredIDCalls())
func (mock *EventRepoMock) GetStartedEventsForTriggeredIDCalls() []struct {
	EventScope models.EventScope
} {
	var calls []struct {
		EventScope models.EventScope
	}
	mock.lockGetStartedEventsForTriggeredID.RLock()
	calls = mock.calls.GetStartedEventsForTriggeredID
	mock.lockGetStartedEventsForTriggeredID.RUnlock()
	return calls
}

// GetTaskSequenceTriggeredEvent calls GetTaskSequenceTriggeredEventFunc.
func (mock *EventRepoMock) GetTaskSequenceTriggeredEvent(eventScope models.EventScope, taskSequenceName string) (*apimodels.KeptnContextExtendedCE, error) {
	if mock.GetTaskSequenceTriggeredEventFunc == nil {
		panic("EventRepoMock.GetTaskSequenceTriggeredEventFunc: method is nil but EventRepo.GetTaskSequenceTriggeredEvent was just called")
	}
	callInfo := struct {
		EventScope       models.EventScope
		TaskSequenceName string
	}{
		EventScope:       eventScope,
		TaskSequenceName: taskSequenceName,
	}
	mock.lockGetTaskSequenceTriggeredEvent.Lock()
	mock.calls.GetTaskSequenceTriggeredEvent = append(mock.calls.GetTaskSequenceTriggeredEvent, callInfo)
	mock.lockGetTaskSequenceTriggeredEvent.Unlock()
	return mock.GetTaskSequenceTriggeredEventFunc(eventScope, taskSequenceName)
}

// GetTaskSequenceTriggeredEventCalls gets all the calls that were made to GetTaskSequenceTriggeredEvent.
// Check the length with:
//     len(mockedEventRepo.GetTaskSequenceTriggeredEventCalls())
func (mock *EventRepoMock) GetTaskSequenceTriggeredEventCalls() []struct {
	EventScope       models.EventScope
	TaskSequenceName string
} {
	var calls []struct {
		EventScope       models.EventScope
		TaskSequenceName string
	}
	mock.lockGetTaskSequenceTriggeredEvent.RLock()
	calls = mock.calls.GetTaskSequenceTriggeredEvent
	mock.lockGetTaskSequenceTriggeredEvent.RUnlock()
	return calls
}

// InsertEvent calls InsertEventFunc.
func (mock *EventRepoMock) InsertEvent(project string, event apimodels.KeptnContextExtendedCE, status common.EventStatus) error {
	if mock.InsertEventFunc == nil {
		panic("EventRepoMock.InsertEventFunc: method is nil but EventRepo.InsertEvent was just called")
	}
	callInfo := struct {
		Project string
		Event   apimodels.KeptnContextExtendedCE
		Status  common.EventStatus
	}{
		Project: project,
		Event:   event,
		Status:  status,
	}
	mock.lockInsertEvent.Lock()
	mock.calls.InsertEvent = append(mock.calls.InsertEvent, callInfo)
	mock.lockInsertEvent.Unlock()
	return mock.InsertEventFunc(project, event, status)
}

// InsertEventCalls gets all the calls that were made to InsertEvent.
// Check the length with:
//     len(mockedEventRepo.InsertEventCalls())
func (mock *EventRepoMock) InsertEventCalls() []struct {
	Project string
	Event   apimodels.KeptnContextExtendedCE
	Status  common.EventStatus
} {
	var calls []struct {
		Project string
		Event   apimodels.KeptnContextExtendedCE
		Status  common.EventStatus
	}
	mock.lockInsertEvent.RLock()
	calls = mock.calls.InsertEvent
	mock.lockInsertEvent.RUnlock()
	return calls
}
