package commands

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"gitlab.com/norzion/temp0/command-handler/app"
	"gitlab.com/norzion/temp0/command-handler/errors"
)

func TestNewSetAlarms(t *testing.T) {
	ds := NewMockstore()

	app := &app.App{
		Log:       log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
		Datastore: ds,
	}

	tests := []struct {
		name string
		want *SetAlarms
	}{
		{
			"NewSetAlarms",
			&SetAlarms{app: app},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSetAlarms(app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetAlarms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetAlarms_Handle(t *testing.T) {
	ds := NewMockstore()
	app := &app.App{
		Log:       log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
		Datastore: ds,
	}

	type args struct {
		ctx         context.Context
		body        []byte
		marshalBody bool
	}
	tests := []struct {
		name    string
		e       *SetAlarms
		args    args
		wantErr bool
	}{
		{
			name: "Legal body",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"testID1": {
								"alarmState": "on",
								"barometer": {
									"low": 15,
									"high": 20
								},
								"temperature": {
									"low": 15,
									"high": 20
								}
							},
							"testID2": {
								"alarmState": "off",
								"barometer": {
									"low": 15,
									"high": 20
								},
								"temperature": {
									"low": 15,
									"high": 20
								}						
							}
						}
					}`,
				),
			},
			wantErr: false,
		},
		{
			name: "Illegal body - empty devices key",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"": {
							"testID1": {
								"alarmState": "on",
								"barometer": {
									"low": 15,
									"high": 20
								},
								"temperature": {
									"low": 15,
									"high": 20
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - empty device ID",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"": {
								"alarmState": "on",
								"barometer": {
									"low": 15,
									"high": 20
								},
								"temperature": {
									"low": 15,
									"high": 20
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - no alarms for device",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"123456": {
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - incorrect alarmState",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"123456": {
								"alarmState": "flaf",
								"barometer": {
									"low": 15,
									"high": 20
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - incorrect barometer bounds",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"123456": {
								"alarmState": "on",
								"barometer": {
									"low": 20,
									"high": 15
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - incorrect temperature bounds",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"123456": {
								"alarmState": "on",
								"temperature": {
									"low": 20,
									"high": 15
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body content",
			e:    NewSetAlarms(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{"flaf": "giraf"}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body ",
			e:    NewSetAlarms(app),
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
				t.Errorf("SetAlarms.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetAlarms_validate(t *testing.T) {
	type args struct {
		body *SetAlarmsBody
	}
	tests := []struct {
		name string
		cmd  *SetAlarms
		args args
		want *errors.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cmd.validate(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAlarms.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
