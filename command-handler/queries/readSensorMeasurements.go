package queries

import (
	"context"
	"encoding/json"

	"github.com/TS-DIY/senz0/command-handler/app"
	"github.com/TS-DIY/senz0/command-handler/errors"
)

// ReadSensorMeasurementsBody ...
type ReadSensorMeasurementsBody struct {
	Devices []string `json:"devices"`
}

// ReadSensorMeasurements ..
type ReadSensorMeasurements struct {
	app *app.App
}

// NewReadSensorMeasurements ..
func NewReadSensorMeasurements(app *app.App) *ReadSensorMeasurements {
	return &ReadSensorMeasurements{app: app}
}

// Handle is a demo Handle for a command..
func (query *ReadSensorMeasurements) Handle(ctx context.Context, b []byte) ([]byte, *errors.Error) {

	// populate our command body using appropriate unmarshalling
	body := &ReadSensorMeasurementsBody{}

	/* if using JSON */
	if err := json.Unmarshal(b, &body); err != nil {
		return nil, errors.NewError("could not decode payload").WithCode(errors.ErrInvalidRequest)
	}

	if err := query.validate(body); err != nil {
		return nil, err
	}

	measurements, err := query.app.Datastore.ReadSensorMeasurements(body.Devices)
	if err != nil {
		return nil, err
	}

	response, stderr := json.Marshal(measurements)
	if stderr != nil {
		return nil, errors.NewError("Could not encode response body").WithCode(errors.ErrServerInternal)
	}

	return response, nil
}

// validate verifies that the command body has valid content
func (query *ReadSensorMeasurements) validate(body *ReadSensorMeasurementsBody) *errors.Error {
	if body.Devices == nil || len(body.Devices) == 0 {
		return errors.NewError("No devices specified in request").WithCode(errors.ErrInvalidRequest)
	}

	return nil
}
