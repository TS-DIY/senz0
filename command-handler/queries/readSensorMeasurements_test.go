package queries

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/TS-DIY/senz0/command-handler/app"
	"github.com/TS-DIY/senz0/command-handler/datastore"
	"github.com/TS-DIY/senz0/command-handler/errors"
)

func TestReadSensorMeasurements_Handle(t *testing.T) {
	ms := NewMockStore()

	ms.Devices["testDevice1"] = &datastore.Temp0Device{
		Measurements: []*datastore.Measure{&datastore.Measure{
			TimeMeasured: 5,
			Temperature:  10,
			Pressure:     1000,
		}},
	}

	type args struct {
		ctx         context.Context
		requestBody []byte
	}

	app := &app.App{
		Datastore: ms,
	}

	tests := []struct {
		name         string
		query        *ReadSensorMeasurements
		args         args
		responseBody []byte
		wantErr      bool
	}{
		{
			name:  "Legal body + working response",
			query: NewReadSensorMeasurements(app),
			args: args{
				ctx: context.Background(),
				requestBody: []byte(
					`{
						"devices": ["testDevice1"]
					}`,
				),
			},
			responseBody: []byte(
				`{"testDevice1":[{"timeMeasured":5,"temperature":10,"pressure":1000}]}`,
			),
			wantErr: false,
		},
		{
			name:  "Illegal request - non-valid json",
			query: NewReadSensorMeasurements(app),
			args: args{
				ctx: context.Background(),
				requestBody: []byte(
					`{
						"devices":
					}`,
				),
			},
			responseBody: nil,
			wantErr:      true,
		},
		{
			name:  "Illegal request - empty device list",
			query: NewReadSensorMeasurements(app),
			args: args{
				ctx: context.Background(),
				requestBody: []byte(
					`{
						"devices":[]
					}`,
				),
			},
			responseBody: nil,
			wantErr:      true,
		},
		{
			name:  "Illegal request - device not registered",
			query: NewReadSensorMeasurements(app),
			args: args{
				ctx: context.Background(),
				requestBody: []byte(
					`{
						"devices":["testDevice2"]
					}`,
				),
			},
			responseBody: nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			responseBody, err := tt.query.Handle(tt.args.ctx, tt.args.requestBody)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSensorMeasurements.Handle() got1 = %v, want %v", err, tt.wantErr)
			}

			// no error, we can now evaluate response bodies
			if err == nil && bytes.Compare(responseBody, tt.responseBody) != 0 {
				t.Errorf("ReadSensorMeasurements.Handle() got = %s, want %s", string(responseBody), string(tt.responseBody))
			}

		})
	}
}

func TestReadSensorMeasurements_validate(t *testing.T) {
	type args struct {
		body *ReadSensorMeasurementsBody
	}
	tests := []struct {
		name  string
		query *ReadSensorMeasurements
		args  args
		want  *errors.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.query.validate(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadSensorMeasurements.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
