package datastore

import "github.com/TS-DIY/senz0/command-handler/errors"

type config struct {
	connection string
}

// Datastore provides a generic interface for instantiating a datastore using any appropriate database driver
type Datastore interface {
	StoreSensorMeasurements(deviceID string, measurements []*Measure) *errors.Error
	SetAlarms(devices map[string]*DeviceAlarm) *errors.Error
	SetAlarmNotification(devices map[string]*AlarmNotification) *errors.Error
	ReadSensorMeasurements(deviceIDs []string) (map[string][]*Measure, *errors.Error)
}

// Temp0Device contains a timeseries of temperature and pressure measurements for a single sensor
type Temp0Device struct {
	Measurements      []*Measure
	AlarmNotification *AlarmNotification
	DeviceAlarm       *DeviceAlarm
}

// Measure contains a single measurement point in time for a single sensor
type Measure struct {
	TimeMeasured int64   `json:"timeMeasured"` // Unix time in nano seconds since epoch
	Temperature  float64 `json:"temperature"`
	Pressure     float64 `json:"pressure"`
}

// DeviceAlarm ..
type DeviceAlarm struct {
	State       string
	Temperature *SensorAlarm
	Barometer   *SensorAlarm
}

// SensorAlarm ..
type SensorAlarm struct {
	Low  float64
	High float64
}

// State ..
type State string

// AlarmNotification ...
type AlarmNotification struct {
	Emails map[string]State
}
