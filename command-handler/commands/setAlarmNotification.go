package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/TS-DIY/senz0/command-handler/app"
	"github.com/TS-DIY/senz0/command-handler/datastore"
	"github.com/TS-DIY/senz0/command-handler/errors"
)

// SetAlarmNotificationBody ...
type SetAlarmNotificationBody struct {
	Devices map[string]*AlarmNotification `json:"devices"`
}

// State defines the state of an alarm
type State string

// AlarmNotification ...
type AlarmNotification struct {
	Emails map[string]State `json:"email"`
}

// SetAlarmNotification ..
type SetAlarmNotification struct {
	app *app.App
}

// NewSetAlarmNotification ..
func NewSetAlarmNotification(app *app.App) *SetAlarmNotification {
	return &SetAlarmNotification{app: app}
}

// Handle is a demo Handle for a command..
func (cmd *SetAlarmNotification) Handle(ctx context.Context, b []byte) *errors.Error {

	// populate our command body using appropriate unmarshalling
	body := &SetAlarmNotificationBody{}

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
func (cmd *SetAlarmNotification) validate(body *SetAlarmNotificationBody) *errors.Error {
	if body.Devices == nil {
		return errors.NewError("No or incorrect request body").WithCode(errors.ErrInvalidRequest)
	}

	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	for id, devices := range body.Devices {
		if id == "" {
			return errors.NewError("Empty temp0 device id").WithCode(errors.ErrInvalidRequest)
		}
		for email, state := range devices.Emails {
			if email == "" {
				return errors.NewError("Empty email").WithCode(errors.ErrInvalidRequest)
			}

			if len(email) > 254 || !rxEmail.MatchString(email) {
				return errors.NewError(fmt.Sprintf("[%s] is not a valid email", email)).WithCode(errors.ErrInvalidRequest)
			}

			if state != "on" && state != "off" {
				return errors.NewError(fmt.Sprintf("[%s] invalid state: %s", email, state)).WithCode(errors.ErrInvalidRequest)
			}
		}
	}

	return nil
}

func (cmd *SetAlarmNotification) store(body *SetAlarmNotificationBody) *errors.Error {
	// convert to datastore format and store
	devices := make(map[string]*datastore.AlarmNotification)
	for id, notifications := range body.Devices {
		alarmNotification := &datastore.AlarmNotification{
			Emails: make(map[string]datastore.State),
		}
		for emails, state := range notifications.Emails {
			alarmNotification.Emails[emails] = datastore.State(state)
		}
		devices[id] = alarmNotification
	}
	return cmd.app.Datastore.SetAlarmNotification(devices)
}
