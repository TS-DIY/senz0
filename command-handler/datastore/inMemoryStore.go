package datastore

import (
	"fmt"

	"gitlab.com/norzion/temp0/command-handler/errors"
)

// InMemoryStore provides a way to run the service without actual database dependencies
type InMemoryStore struct {
	devices map[string]*Temp0Device
}

// NewInMemoryStore creates a postgres database based datastore using "URI" and "maxConn" to connect
func NewInMemoryStore() *InMemoryStore {
	ds := &InMemoryStore{}

	ds.devices = make(map[string]*Temp0Device)

	return ds
}

// NewTemp0Device is used to register a new device the first time it broadcasts measurements
func NewTemp0Device() *Temp0Device {
	return &Temp0Device{
		Measurements: []*Measure{},
		AlarmNotification: &AlarmNotification{
			Emails: make(map[string]State),
		},
		DeviceAlarm: &DeviceAlarm{
			State:       "off",
			Temperature: &SensorAlarm{},
			Barometer:   &SensorAlarm{},
		},
	}
}

// StoreSensorMeasurements ..
func (ds *InMemoryStore) StoreSensorMeasurements(deviceID string, measurements []*Measure) *errors.Error {
	// if no device previously registered, then create (failsafe)
	if ds.devices[deviceID] == nil {
		ds.devices[deviceID] = NewTemp0Device()
	}

	// store the measurements
	ds.devices[deviceID].Measurements = append(ds.devices[deviceID].Measurements, measurements...)

	return nil
}

// SetAlarms ..
func (ds *InMemoryStore) SetAlarms(devices map[string]*DeviceAlarm) *errors.Error {
	// device must be registered (alive) before alarms are set
	for id := range devices {
		if ds.devices[id] == nil {
			return errors.NewError(fmt.Sprintf("[%s]: Invalid device", id)).WithCode(errors.ErrInvalidRequest)
		}
	}

	// store the alarm
	for id, alarm := range devices {
		ds.devices[id].DeviceAlarm.State = alarm.State
		if alarm.Barometer != nil {
			ds.devices[id].DeviceAlarm.Barometer = alarm.Barometer
		}
		if alarm.Temperature != nil {
			ds.devices[id].DeviceAlarm.Temperature = alarm.Temperature
		}
	}

	return nil
}

// SetAlarmNotification ..
func (ds *InMemoryStore) SetAlarmNotification(devices map[string]*AlarmNotification) *errors.Error {
	// device must be registered (alive) before alarms are set
	for id := range devices {
		if ds.devices[id] == nil {
			return errors.NewError(fmt.Sprintf("[%s]: Invalid device", id)).WithCode(errors.ErrInvalidRequest)
		}
	}

	// store all notification receivers
	for id, notification := range devices {
		for email, state := range notification.Emails {
			ds.devices[id].AlarmNotification.Emails[email] = state
		}
	}
	return nil
}

// ReadSensorMeasurements ..
func (ds *InMemoryStore) ReadSensorMeasurements(deviceIDs []string) (map[string][]*Measure, *errors.Error) {

	devices := make(map[string][]*Measure)
	for _, id := range deviceIDs {
		if ds.devices[id] == nil {
			return nil, errors.NewError(fmt.Sprintf("[%s]: Invalid device", id)).WithCode(errors.ErrInvalidRequest)
		}
		devices[id] = ds.devices[id].Measurements
	}
	return devices, nil

}
