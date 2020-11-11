package cmd

import (
	v8 "github.com/v8platform/v8"
	"reflect"
	"testing"
)

func Test_infobaseFromString(t *testing.T) {
	type args struct {
		connectionString string
	}
	tests := []struct {
		name string
		args args
		want v8.Infobase
	}{
		{"file relative",
			args{connectionString: "/F./foo"},
			v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "./foo",
				Locale:   "",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := infobaseFromString(tt.args.connectionString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infobaseFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
