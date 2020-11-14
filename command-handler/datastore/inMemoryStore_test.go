package datastore

import (
	"testing"
)

func TestInMemoryStore_StoreSensorMeasurements(t *testing.T) {
	type args struct {
		deviceID     string
		measurements []*Measure
	}
	tests := []struct {
		name    string
		ds      *InMemoryStore
		args    args
		wantErr bool
	}{
		{
			name: "Legal",
			ds:   NewInMemoryStore(),
			args: args{
				deviceID:     "testID1",
				measurements: []*Measure{&Measure{}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ds.StoreSensorMeasurements(tt.args.deviceID, tt.args.measurements); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.StoreSensorMeasurements(): error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryStore_SetAlarms(t *testing.T) {
	type args struct {
		devices map[string]*DeviceAlarm
	}

	test1DS := NewInMemoryStore()
	test1DS.devices["deviceID1"] = NewTemp0Device()

	test1DeviceAlarms := make(map[string]*DeviceAlarm)
	test1DeviceAlarms["deviceID1"] = &DeviceAlarm{
		State:       "On",
		Barometer:   &SensorAlarm{},
		Temperature: &SensorAlarm{},
	}
	test2DeviceAlarms := make(map[string]*DeviceAlarm)
	test2DeviceAlarms["deviceID2"] = &DeviceAlarm{
		State:       "On",
		Barometer:   &SensorAlarm{},
		Temperature: &SensorAlarm{},
	}

	tests := []struct {
		name    string
		ds      *InMemoryStore
		args    args
		wantErr bool
	}{
		{
			name: "Legal",
			ds:   test1DS,
			args: args{
				devices: test1DeviceAlarms,
			},
			wantErr: false,
		},
		{
			name: "Illegal - no device registered",
			ds:   test1DS,
			args: args{
				devices: test2DeviceAlarms,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ds.SetAlarms(tt.args.devices); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.SetAlarms() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryStore_SetAlarmNotification(t *testing.T) {
	type args struct {
		devices map[string]*AlarmNotification
	}

	test1DS := NewInMemoryStore()
	test1DS.devices["deviceID1"] = NewTemp0Device()

	testDevice1Notifications := make(map[string]*AlarmNotification)
	testDevice1Notifications["deviceID1"] = &AlarmNotification{
		Emails: make(map[string]State),
	}
	testDevice1Notifications["deviceID1"].Emails["flaf@flaf.de"] = "On"

	testDevice2Notifications := make(map[string]*AlarmNotification)
	testDevice2Notifications["deviceID2"] = &AlarmNotification{
		Emails: make(map[string]State),
	}
	testDevice2Notifications["deviceID2"].Emails["flaf@flaf.de"] = "On"

	tests := []struct {
		name    string
		ds      *InMemoryStore
		args    args
		wantErr bool
	}{
		{
			name: "Legal",
			ds:   test1DS,
			args: args{
				devices: testDevice1Notifications,
			},
			wantErr: false,
		},
		{
			name: "Illegal - no device registered",
			ds:   test1DS,
			args: args{
				devices: testDevice2Notifications,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ds.SetAlarmNotification(tt.args.devices); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.SetAlarmNotification() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryStore_ReadSensorMeasurements(t *testing.T) {
	type args struct {
		deviceIDs []string
	}

	test1DS := NewInMemoryStore()
	test1DS.devices["testDevice1"] = NewTemp0Device()

	tests := []struct {
		name    string
		ds      *InMemoryStore
		args    args
		wantErr bool
	}{
		{
			name: "Legal",
			ds:   test1DS,
			args: args{
				deviceIDs: []string{"testDevice1"},
			},
			wantErr: false,
		},
		{
			name: "Illegal - no device registered",
			ds:   test1DS,
			args: args{
				deviceIDs: []string{"testDevice2"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				if _, err := tt.ds.ReadSensorMeasurements(tt.args.deviceIDs); (err != nil) != tt.wantErr {
					t.Errorf("InMemoryStore.ReadSensorMeasurements() = %v, want %v", err, tt.wantErr)
				}
			})
		})
	}
}
