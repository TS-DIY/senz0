package datastore

// NB: database interactions require integration test setup with actual db

// import (
// 	"testing"
// )

// func TestNewPSQLStore(t *testing.T) {
// 	type args struct {
// 		URI     string
// 		maxConn int
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "Database connection",
// 			args: args{
// 				URI:     "postgres://norzion_test:flafgiraf@localhost/temp0_test",
// 				maxConn: 5},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Database connection, error in URI",
// 			args: args{
// 				URI:     "postgres://norzion_test:flafgiraf@:localhost/temp0_test",
// 				maxConn: 5},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if _, gotErr := NewPSQLStore(tt.args.URI, tt.args.maxConn); gotErr != nil && !tt.wantErr || gotErr == nil && tt.wantErr {
// 				t.Errorf("NewPSQLStore(): returned %v", gotErr)
// 			}
// 		})
// 	}
// }
