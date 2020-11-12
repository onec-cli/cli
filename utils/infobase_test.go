package utils_test

import (
	"github.com/onec-cli/cli/utils"
	v8 "github.com/v8platform/v8"
	"reflect"
	"testing"
)

func TestGetInfobase(t *testing.T) {
	//if testing.Short() {
	//	t.Skip("skipping test in short mode.")
	//}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    v8.Infobase
		wantErr bool
	}{
		{
			"Empty",
			args{s: ""},
			nil,
			true,
		},
		{
			"Spaces and Cyrillic",
			args{s: " /F./f_Фл_o "},
			&v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "./f_Фл_o",
				Locale:   "",
			},
			false,
		},
		{
			"Prefix UPPER vs lower",
			args{s: "/f./foo"},
			&v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "./foo",
				Locale:   "",
			},
			false,
		},
		{
			"File relative path",
			args{s: "/F./foo"},
			&v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "./foo",
				Locale:   "",
			},
			false,
		},
		{
			"File=",
			args{s: "File=\"C:\\foo\\boo\";"},
			&v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     `C:\foo\boo`,
				Locale:   "",
			},
			false,
		},
		{
			"Server invalid string",
			args{s: "/Sfoo"},
			nil,
			true,
		},
		{
			"Server sep",
			args{s: "/Sfoo\\boo"},
			&v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "foo",
				Ref:      "boo",
			},
			false,
		},
		{
			"Server tcp",
			args{s: "/Stcp://foo:1641/boo"},
			&v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "tcp://foo:1641",
				Ref:      "boo",
			},
			false,
		},
		{
			"Server IPv6",
			args{s: "/S[fe10::c47b:90b7:fa32:a2fa%12]/boo"},
			&v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "[fe10::c47b:90b7:fa32:a2fa%12]",
				Ref:      "boo",
			},
			false,
		},
		{
			"Server multi-claster",
			args{s: "/S127.0.0.1:1541,127.0.0.2:1542/boo"},
			&v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "127.0.0.1:1541,127.0.0.2:1542",
				Ref:      "boo",
			},
			false,
		},
		{
			"Server",
			args{s: "/S127.0.0.1:1541/boo"},
			&v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "127.0.0.1:1541",
				Ref:      "boo",
			},
			false,
		},
		{
			"Srvr= Ref= ",
			args{s: "Srvr=\"foo\";Ref=\"boo\";"},
			&v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "foo",
				Ref:      "boo",
			},
			false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.NewInfobase(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInfobase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInfobase() got = %v, want %v", got, tt.want)
			}
		})

	}
}
