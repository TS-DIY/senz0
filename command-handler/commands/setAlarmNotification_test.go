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

func TestNewSetAlarmNotification(t *testing.T) {
	ds := NewMockstore()

	app := &app.App{
		Log:       log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
		Datastore: ds,
	}

	tests := []struct {
		name string
		want *SetAlarmNotification
	}{
		{
			"NewSetAlarmNotification",
			&SetAlarmNotification{app: app},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSetAlarmNotification(app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetAlarmNotification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetAlarmNotification_Handle(t *testing.T) {
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
		e       *SetAlarmNotification
		args    args
		wantErr bool
	}{
		{
			name: "Legal body",
			e:    NewSetAlarmNotification(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"testID1": {
								"email": {
									"flaf@giraf.dk": "on",
									"boing@hat.dk": "off"
								}
							}
						}
					}`,
				),
			},
			wantErr: false,
		},
		{
			name: "Illegal body - empty device id",
			e:    NewSetAlarmNotification(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"": {
								"email": {
									"flaf@giraf.dk": "on"
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - empty email",
			e:    NewSetAlarmNotification(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"123456": {
								"email": {
									"": "on"
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - invalid email",
			e:    NewSetAlarmNotification(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"123456": {
								"email": {
									"flaf@@moo": "on"
								}
							}
						}
					}`,
				),
			},
			wantErr: true,
		},
		{
			name: "Illegal body - invalid email",
			e:    NewSetAlarmNotification(app),
			args: args{
				ctx: context.Background(),
				body: []byte(
					`{
						"devices": {
							"123456": {
								"email": {
									"flaf@giraf.dk": "nope"
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
			e:    NewSetAlarmNotification(app),
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
			e:    NewSetAlarmNotification(app),
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
				t.Errorf("SetAlarmNotification.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetAlarmNotification_validate(t *testing.T) {
	type args struct {
		body *SetAlarmNotificationBody
	}
	tests := []struct {
		name string
		cmd  *SetAlarmNotification
		args args
		want *errors.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cmd.validate(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAlarmNotification.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
