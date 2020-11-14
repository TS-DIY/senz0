package queries

import (
	"gitlab.com/norzion/temp0/command-handler/datastore"
	"gitlab.com/norzion/temp0/command-handler/errors"
)

type MockStore struct {
	Test    string
	Devices map[string]*datastore.Temp0Device
}

func NewMockStore() *MockStore {
	ms := &MockStore{
		Devices: make(map[string]*datastore.Temp0Device),
		Test:    "valid",
	}

	return ms
}

func (ms *MockStore) StoreSensorMeasurements(deviceID string, measurements []*datastore.Measure) *errors.Error {
	return nil
}
func (ms *MockStore) SetAlarms(devices map[string]*datastore.DeviceAlarm) *errors.Error {
	return nil
}
func (ms *MockStore) SetAlarmNotification(devices map[string]*datastore.AlarmNotification) *errors.Error {
	return nil
}
func (ms *MockStore) ReadSensorMeasurements(deviceIDs []string) (map[string][]*datastore.Measure, *errors.Error) {

	switch ms.Test {
	case "valid":
		devices := make(map[string][]*datastore.Measure)
		for _, id := range deviceIDs {
			if ms.Devices[id] == nil {
				return nil, errors.NewError("ID does not exist").WithCode(errors.ErrInvalidRequest)
			}
			devices[id] = ms.Devices[id].Measurements
		}
		return devices, nil
	}
	return nil, nil
}
