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
		connectionString string
	}
	tests := []struct {
		name string
		args args
		want v8.Infobase
	}{
		{
			"Empty",
			args{connectionString: ""},
			nil,
		},
		{
			"Spaces and Cyrillic",
			args{connectionString: " /F./f_Фл_o "},
			v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "./f_Фл_o",
				Locale:   "",
			},
		},
		{
			"Prefix UPPER vs lower",
			args{connectionString: "/f./foo"},
			v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "./foo",
				Locale:   "",
			},
		},
		{
			"File relative path",
			args{connectionString: "/F./foo"},
			v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "./foo",
				Locale:   "",
			},
		},
		{
			"File=",
			args{connectionString: "File=\"C:\\foo\\boo\";"},
			v8.FileInfoBase{
				InfoBase: v8.InfoBase{},
				File:     "C:\\foo\\boo",
				Locale:   "",
			},
		},
		{
			"Server sep",
			args{connectionString: "/Sfoo\\boo"},
			v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "foo",
				Ref:      "boo",
			},
		},
		{
			"Server bad sep",
			args{connectionString: "/Sfoo\\boo/fff"},
			nil,
		},
		{
			"Server tcp",
			args{connectionString: "/Stcp://foo:1641/boo"},
			v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "tcp://foo:1641",
				Ref:      "boo",
			},
		},
		{
			"Server IPv6",
			args{connectionString: "/S[fe10::c47b:90b7:fa32:a2fa%12]/boo"},
			v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "[fe10::c47b:90b7:fa32:a2fa%12]",
				Ref:      "boo",
			},
		},
		{
			"Server multi-claster",
			args{connectionString: "/S127.0.0.1:1541,127.0.0.2:1542/boo"},
			v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "127.0.0.1:1541,127.0.0.2:1542",
				Ref:      "boo",
			},
		},
		{
			"Server",
			args{connectionString: "/S127.0.0.1:1541/boo"},
			v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "127.0.0.1:1541",
				Ref:      "boo",
			},
		},
		{
			"Srvr= Ref= ",
			args{connectionString: "Srvr=\"foo\";Ref=\"boo\";"},
			v8.ServerInfoBase{
				InfoBase: v8.InfoBase{},
				Srvr:     "foo",
				Ref:      "boo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetInfobase(tt.args.connectionString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfobase() = %v, want %v", got, tt.want)
			}
		})
	}
}
