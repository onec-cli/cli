package platform_test

import (
	"github.com/onec-cli/cli/internal/platform"
	"reflect"
	"testing"
)

func TestGetDefaultOptions(t *testing.T) {
	type args struct {
		opts map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				opts: map[string]interface{}{
					"dbsrvr":  "db",
					"db":      "ib",
					"dbuid":   "postgres",
					"crsqldb": true,
					"usr":     "",
					"pwd":     nil,
					"dbms":    "PostgreSQL",
				},
			},
			want: []string{
				"DBMS=PostgreSQL",
				"DBSrvr=db",
				"DB=ib",
				"DBUID=postgres",
				"CrSQLDB=Y",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := platform.DefaultOptions(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultOptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
