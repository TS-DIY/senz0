package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TS-DIY/senz0/command-handler/app"
	"github.com/TS-DIY/senz0/command-handler/datastore"
	"github.com/TS-DIY/senz0/command-handler/errors"
)

// SetAlarmsBody ...
type SetAlarmsBody struct {
	Devices map[string]*Temp0Device `json:"devices"`
}

// Temp0Device represents a single temp0 device, represented by an ID with several sensors on it
type Temp0Device struct {
	AlarmState  string       `json:"alarmState"`
	Barometer   *SensorAlarm `json:"barometer"`
	Temperature *SensorAlarm `json:"temperature"`
}

// SensorAlarm sets an alarm for a specific sensor. Equal values disables the alarm.
type SensorAlarm struct {
	// State int     `json:"state"`
	Low  float64 `json:"low"`  // omitting in request defaults to 0
	High float64 `json:"high"` // omitting in request defaults to 0
}

// SetAlarms ..
type SetAlarms struct {
	app *app.App
}

// NewSetAlarms ..
func NewSetAlarms(app *app.App) *SetAlarms {
	return &SetAlarms{app: app}
}

// Handle is a demo Handle for a command..
func (cmd *SetAlarms) Handle(ctx context.Context, b []byte) *errors.Error {

	// populate our command body using appropriate unmarshalling
	body := &SetAlarmsBody{}

	/* if using JSON */
	if err := json.Unmarshal(b, &body); err != nil {
		return errors.NewError("could not decode payload").WithCode(errors.ErrInvalidRequest)
	}

	if err := cmd.validate(body); err != nil {
		return err
	}

	return cmd.store(body)
}

// validate verifies that the command body has valid content
func (cmd *SetAlarms) validate(body *SetAlarmsBody) *errors.Error {
	if body.Devices == nil {
		return errors.NewError("No or incorrect request body").WithCode(errors.ErrInvalidRequest)
	}

	for k, v := range body.Devices {
		if k == "" {
			return errors.NewError("Empty temp0 device id").WithCode(errors.ErrInvalidRequest)
		}
		if v != nil && v.Barometer == nil && v.Temperature == nil {
			return errors.NewError(fmt.Sprintf("[%s] No alarms specified", k)).WithCode(errors.ErrInvalidRequest)
		}

		if v != nil && v.AlarmState != "on" && v.AlarmState != "off" {
			return errors.NewError(fmt.Sprintf("[%s] Barometer alarm has incorrect state: %s", k, v.AlarmState)).WithCode(errors.ErrInvalidRequest)

		}

		if v != nil && v.Barometer != nil && v.Barometer.Low > v.Barometer.High {
			return errors.NewError(fmt.Sprintf("[%s] Barometer alarm has incorrect bounds", k)).WithCode(errors.ErrInvalidRequest)
		}

		if v != nil && v.Temperature != nil && v.Temperature.Low > v.Temperature.High {
			return errors.NewError(fmt.Sprintf("[%s] Temperature alarm has incorrect bounds", k)).WithCode(errors.ErrInvalidRequest)
		}
	}

	return nil
}

func (cmd *SetAlarms) store(body *SetAlarmsBody) *errors.Error {
	devices := make(map[string]*datastore.DeviceAlarm)
	for id, alarm := range body.Devices {
		dsAlarm := &datastore.DeviceAlarm{}
		if alarm.Barometer != nil {
			dsAlarm.Barometer = &datastore.SensorAlarm{
				Low:  alarm.Barometer.Low,
				High: alarm.Barometer.High,
			}
		}
		if alarm.Temperature != nil {
			dsAlarm.Temperature = &datastore.SensorAlarm{
				Low:  alarm.Temperature.Low,
				High: alarm.Temperature.High,
			}
		}
		devices[id] = dsAlarm
	}

	return cmd.app.Datastore.SetAlarms(devices)
}
