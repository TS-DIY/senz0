package commands

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/TS-DIY/senz0/command-handler/app"
	"github.com/TS-DIY/senz0/command-handler/errors"
)

func TestNewReportSensorMeasurements(t *testing.T) {
	ds := NewMockstore()

	app := &app.App{
		Log:       log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
		Datastore: ds,
	}

	tests := []struct {
		name string
		want *ReportSensorMeasurements
	}{
		{
			"NewReportSensorMeasurements",
			&ReportSensorMeasurements{app: app},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReportSensorMeasurements(app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReportSensorMeasurements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReportSensorMeasurements_Handle(t *testing.T) {
	ds := NewMockstore()
	app := &app.App{
		Log:       log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
		Datastore: ds,
	}

	type args struct {
		ctx  context.Context
		body []byte
	}
	tests := []struct {
		name    string
		e       *ReportSensorMeasurements
		args    args
		wantErr bool
	}{
		{
			name: "Legal body",
			e:    NewReportSensorMeasurements(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"temp0ID": "testDeviceID",
						"measurements": [{
							"timeMeasured": 1541763380,
							"temperature": 25.0,
							"pressure": 1000.0
						}]
					}`,
				),
			},
			wantErr: false,
		},
		{
			name: "Illegal body - no device ID",
			e:    NewReportSensorMeasurements(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"temp0ID": "",
						"measurements": [{
							"timeMeasured": 1541763380,
							"temperature": 25.0,
							"pressure": 1000.0
						}]
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - no measurements",
			e:    NewReportSensorMeasurements(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"temp0ID": "1234",
						"measurements": []
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body content",
			e:    NewReportSensorMeasurements(app),
			args: args{
				ctx:  context.Background(),
				body: []byte("{someVariable: hat}"),
			},
			wantErr: true,
		},
		{
			name: "Illegal body ",
			e:    NewReportSensorMeasurements(app),
			args: args{
				ctx:  context.Background(),
				body: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Handle(tt.args.ctx, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("ReportSensorMeasurements.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReportSensorMeasurements_validate(t *testing.T) {
	type args struct {
		body *ReportSensorMeasurementsBody
	}
	tests := []struct {
		name string
		cmd  *ReportSensorMeasurements
		args args
		want *errors.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cmd.validate(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReportSensorMeasurements.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
