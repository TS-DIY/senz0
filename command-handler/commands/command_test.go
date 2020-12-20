package commands

import (
	"github.com/TS-DIY/senz0/command-handler/datastore"
	"github.com/TS-DIY/senz0/command-handler/errors"
)

type MockStore struct {
}

func NewMockstore() *MockStore {
	return &MockStore{}
}

func (ds *MockStore) StoreSensorMeasurements(deviceID string, measurements []*datastore.Measure) *errors.Error {
	return nil
}

func (ds *MockStore) SetAlarms(devices map[string]*datastore.DeviceAlarm) *errors.Error {
	return nil
}

func (ds *MockStore) SetAlarmNotification(devices map[string]*datastore.AlarmNotification) *errors.Error {
	return nil
}
func (ds *MockStore) ReadSensorMeasurements(deviceIDs []string) (map[string][]*datastore.Measure, *errors.Error) {
	return nil, nil
}
