package api

import (
	"github.com/v8platform/runner"
	"reflect"
	"testing"
)

//todo перенести во внешние тесты
func TestCreateInfobase(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name    string
		args    args
		want    runner.Command
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateInfobase(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateInfobase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInfobase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectionString_Check(t *testing.T) {
	type fields struct {
		connectionString string
		values           []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connectionString{
				connectionString: tt.fields.connectionString,
				values:           tt.fields.values,
			}
			if err := c.Check(); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_connectionString_Command(t *testing.T) {
	type fields struct {
		connectionString string
		values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connectionString{
				connectionString: tt.fields.connectionString,
				values:           tt.fields.values,
			}
			if got := c.Command(); got != tt.want {
				t.Errorf("Command() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectionString_Values(t *testing.T) {
	type fields struct {
		connectionString string
		values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connectionString{
				connectionString: tt.fields.connectionString,
				values:           tt.fields.values,
			}
			if got := c.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectionString_parse(t *testing.T) {
	type fields struct {
		connectionString string
		values           []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "file",
			fields: fields{
				connectionString: "/F./foo",
			},
			want: []string{
				"File=./foo",
			},
			wantErr: false,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connectionString{
				connectionString: tt.fields.connectionString,
				values:           tt.fields.values,
			}
			if err := c.parse(); (err != nil) != tt.wantErr || !reflect.DeepEqual(c.values, tt.want) {
				t.Errorf("parse() error = %v, want %v, wantErr %v", err, tt.want, tt.wantErr)
			}
		})
	}
}
