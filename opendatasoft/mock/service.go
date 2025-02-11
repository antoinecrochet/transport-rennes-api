// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=mock/service.go
//

// Package mock_opendatasoft is a generated GoMock package.
package mock_opendatasoft

import (
	reflect "reflect"

	opendatasoft "github.com/antoinecrochet/transport-rennes-api/opendatasoft"
	gomock "go.uber.org/mock/gomock"
)

// MockOpendatasoftClientInterface is a mock of OpendatasoftClientInterface interface.
type MockOpendatasoftClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOpendatasoftClientInterfaceMockRecorder
	isgomock struct{}
}

// MockOpendatasoftClientInterfaceMockRecorder is the mock recorder for MockOpendatasoftClientInterface.
type MockOpendatasoftClientInterfaceMockRecorder struct {
	mock *MockOpendatasoftClientInterface
}

// NewMockOpendatasoftClientInterface creates a new mock instance.
func NewMockOpendatasoftClientInterface(ctrl *gomock.Controller) *MockOpendatasoftClientInterface {
	mock := &MockOpendatasoftClientInterface{ctrl: ctrl}
	mock.recorder = &MockOpendatasoftClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpendatasoftClientInterface) EXPECT() *MockOpendatasoftClientInterfaceMockRecorder {
	return m.recorder
}

// SearchUpcomingBus mocks base method.
func (m *MockOpendatasoftClientInterface) SearchUpcomingBus(stopName, busLineName, destination string) (*opendatasoft.UpcomingBus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUpcomingBus", stopName, busLineName, destination)
	ret0, _ := ret[0].(*opendatasoft.UpcomingBus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUpcomingBus indicates an expected call of SearchUpcomingBus.
func (mr *MockOpendatasoftClientInterfaceMockRecorder) SearchUpcomingBus(stopName, busLineName, destination any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUpcomingBus", reflect.TypeOf((*MockOpendatasoftClientInterface)(nil).SearchUpcomingBus), stopName, busLineName, destination)
}
