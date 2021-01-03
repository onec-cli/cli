package platform_test

import (
	"github.com/onec-cli/cli/internal/platform"
	"reflect"
	"testing"
)

func TestNewConnStrErrors(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		{
			name: "empty string",
			args: "",
		},
		{
			name: "/S have no sep",
			args: "/Sfoo",
		},
		{
			name: "/S invalid sep",
			args: "/Stcp://foo:1641/boo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := platform.NewConnStr(tt.args)
			if got != nil {
				t.Errorf("NewConnStr() got = %v, want <nil>", got)
				return
			}
			if err == nil {
				t.Error("NewConnStr() there must be an error")
				return
			}
		})
	}
}

func TestNewConnStrSuccess(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "cyrillic",
			args: "/F./f_Фл_o",
			want: []string{
				`File=./f_Фл_o`,
			},
		},
		{
			name: "end delimiter",
			args: "File=./foo;",
			want: []string{
				`File=./foo`,
			},
		},
		{
			name: "missed param",
			args: "File=./foo;;Local=ru",
			want: []string{
				`File=./foo`,
				`Local=ru`,
			},
		},
		{
			name: "case sensitivity",
			args: "FiLE=./foo",
			want: []string{
				`FiLE=./foo`,
			},
		},
		{
			name: "/F",
			args: "/F./foo",
			want: []string{
				`File=./foo`,
			},
		},
		{
			name: "/F spaces",
			args: " /F ./foo ",
			want: []string{
				`File=./foo`,
			},
		},
		{
			name: "/F prefix UPPER vs lower",
			args: "/f./foo",
			want: []string{
				`File=./foo`,
			},
		},
		{
			name: "File= windows",
			args: `File=C:\foo\boo`,
			want: []string{
				`File=C:\foo\boo`,
			},
		},
		{
			name: "File= unix",
			args: "File=/foo/boo",
			want: []string{
				`File=/foo/boo`,
			},
		},
		{
			name: "File=... quotes",
			args: `File="C:\foo\boo"`,
			want: []string{
				`File="C:\foo\boo"`,
			},
		},
		{
			name: "/S",
			args: `/S127.0.0.1:1541\boo`,
			want: []string{
				`Srvr=127.0.0.1:1541`,
				`Ref=boo`,
			},
		},
		{
			name: "/S tcp",
			args: `/Stcp://foo:1641\boo`,
			want: []string{
				`Srvr=tcp://foo:1641`,
				`Ref=boo`,
			},
		},
		{
			name: "/S IPv6",
			args: `/S[fe10::c47b:90b7:fa32:a2fa%12]\boo`,
			want: []string{
				`Srvr=[fe10::c47b:90b7:fa32:a2fa%12]`,
				`Ref=boo`,
			},
		},
		{
			name: "/S multi-claster",
			args: `/S127.0.0.1:1541,127.0.0.2:1542\boo`,
			want: []string{
				`Srvr=127.0.0.1:1541,127.0.0.2:1542`,
				`Ref=boo`,
			},
		},
		{
			name: "Srvr= Ref=",
			args: `Srvr="foo";Ref=boo;`,
			want: []string{
				`Srvr="foo"`,
				`Ref=boo`,
			},
		},
		{
			name: "Ref= Srvr=",
			args: `Ref=boo;Srvr="foo";`,
			want: []string{
				`Ref=boo`,
				`Srvr="foo"`,
			},
		},
		{
			name: "Srvr= Ref= tcp",
			args: `Srvr=tcp://foo:1641;Ref=boo`,
			want: []string{
				`Srvr=tcp://foo:1641`,
				`Ref=boo`,
			},
		},
		{
			name: "Srvr= Ref= IPv6",
			args: `Srvr=[fe10::c47b:90b7:fa32:a2fa%12];Ref=boo`,
			want: []string{
				`Srvr=[fe10::c47b:90b7:fa32:a2fa%12]`,
				`Ref=boo`,
			},
		},
		{
			name: "Srvr= Ref= multi-claster",
			args: `Srvr=127.0.0.1:1541,127.0.0.2:1542;Ref=boo`,
			want: []string{
				`Srvr=127.0.0.1:1541,127.0.0.2:1542`,
				`Ref=boo`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := platform.NewConnStr(tt.args)
			if err != nil {
				t.Errorf("NewConnStr() error = %v, want <nil>", err)
				return
			}
			if !reflect.DeepEqual(got.Values(), tt.want) {
				t.Errorf("NewConnStr() got = %v, want %v", got.Values(), tt.want)
			}
		})
	}
}

//
//func Test_connectionString_Command(t *testing.T) {
//	c := &connectionString{}
//	want := "CREATEINFOBASE"
//	got := c.Command()
//	if got != want {
//		t.Errorf("Command() = %v, want %v", got, want)
//	}
//}
//
//func Test_connectionString_Check(t *testing.T) {
//	c := &connectionString{}
//	got := c.Check()
//	if got != nil {
//		t.Errorf("Command() = %v, want %v", got, nil)
//	}
//}
