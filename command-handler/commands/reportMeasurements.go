package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TS-DIY/senz0/command-handler/app"
	"github.com/TS-DIY/senz0/command-handler/datastore"
	"github.com/TS-DIY/senz0/command-handler/errors"
)

// ReportSensorMeasurementsBody ...
type ReportSensorMeasurementsBody struct {
	Temp0ID      string     `json:"temp0ID"`
	Measurements []*Measure `json:"measurements"`
}

// Measure contains a single measurement point in time for a single sensor
type Measure struct {
	TimeMeasured int64   `json:"timeMeasured"`
	Temperature  float64 `json:"temperature"`
	Pressure     float64 `json:"pressure"`
}

// ReportSensorMeasurements ..
type ReportSensorMeasurements struct {
	app *app.App
}

// NewReportSensorMeasurements ..
func NewReportSensorMeasurements(app *app.App) *ReportSensorMeasurements {
	return &ReportSensorMeasurements{app: app}
}

// Handle is a demo Handle for a command..
func (cmd *ReportSensorMeasurements) Handle(ctx context.Context, b []byte) *errors.Error {

	// populate our command body using appropriate unmarshalling
	body := &ReportSensorMeasurementsBody{}

	/* if using JSON */
	if err := json.Unmarshal(b, &body); err != nil {
		return errors.NewError("could not decode payload").WithCode(errors.ErrInvalidRequest)
	}

	if err := cmd.validate(body); err != nil {
		return err
	}

	// write to database
	measurements := make([]*datastore.Measure, len(body.Measurements))
	for i, val := range body.Measurements {
		measurements[i] = &datastore.Measure{
			TimeMeasured: val.TimeMeasured,
			Temperature:  val.Temperature,
			Pressure:     val.Pressure,
		}
	}

	cmd.app.Datastore.StoreSensorMeasurements(body.Temp0ID, measurements)

	return nil
}

// validate verifies that the command body has valid content
func (cmd *ReportSensorMeasurements) validate(body *ReportSensorMeasurementsBody) *errors.Error {

	if body.Temp0ID == "" {
		return errors.NewError("No temp0 id specified").WithCode(errors.ErrInvalidRequest)
	}

	if body.Measurements == nil || len(body.Measurements) == 0 {
		return errors.NewError(fmt.Sprintf("[%s] No measurements received", body.Temp0ID)).WithCode(errors.ErrInvalidRequest)
	}

	return nil
}
