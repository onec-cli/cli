package api

import (
	"reflect"
	"testing"
)

func Test_connectionString_Command(t *testing.T) {
	c := &connectionString{}
	want := "CREATEINFOBASE"
	got := c.Command()
	if got != want {
		t.Errorf("Command() = %v, want %v", got, want)
	}
}

func Test_connectionString_Check(t *testing.T) {
	c := &connectionString{}
	got := c.Check()
	if got != nil {
		t.Errorf("Command() = %v, want %v", got, nil)
	}
}

func Test_connectionString_Values(t *testing.T) {
	c := &connectionString{
		values: []string{"param1", "param2"},
	}
	want := []string{"param1", "param2"}
	got := c.Values()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Command() = %v, want %v", got, want)
	}
}

func Test_connectionString_removeEmpty(t *testing.T) {
	type fields struct {
		connectionString string
		values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "not init",
			fields: fields{
				values: nil,
			},
			want: nil,
		},
		{
			name: "have empty",
			fields: fields{
				values: []string{
					"",
					"param1",
					"",
					"param2",
					"",
				},
			},
			want: []string{
				"param1",
				"param2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connectionString{
				values: tt.fields.values,
			}
			c.removeEmpty()
			if !reflect.DeepEqual(c.values, tt.want) {
				t.Errorf("removeEmpty() values = %v, want %v", c.values, tt.want)
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
			name: "empty string",
			fields: fields{
				connectionString: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "cyrillic",
			fields: fields{
				connectionString: "/F./f_Фл_o",
			},
			want: []string{
				`File=./f_Фл_o`,
			},
			wantErr: false,
		},
		{
			name: "end delimiter",
			fields: fields{
				connectionString: "File=./foo;",
			},
			want: []string{
				`File=./foo`,
			},
			wantErr: false,
		},
		{
			name: "missed param",
			fields: fields{
				connectionString: "File=./foo;;Local=ru",
			},
			want: []string{
				`File=./foo`,
				`Local=ru`,
			},
			wantErr: false,
		},
		{
			name: "case sensitivity",
			fields: fields{
				connectionString: "FiLE=./foo",
			},
			want: []string{
				`FiLE=./foo`,
			},
			wantErr: false,
		},
		{
			name: "/F",
			fields: fields{
				connectionString: "/F./foo",
			},
			want: []string{
				`File=./foo`,
			},
			wantErr: false,
		},
		{
			name: "/F spaces",
			fields: fields{
				connectionString: " /F ./foo ",
			},
			want: []string{
				`File=./foo`,
			},
			wantErr: false,
		},
		{
			name: "/F prefix UPPER vs lower",
			fields: fields{
				connectionString: "/f./foo",
			},
			want: []string{
				`File=./foo`,
			},
			wantErr: false,
		},
		{
			name: "File= windows",
			fields: fields{
				connectionString: `File=C:\foo\boo`,
			},
			want: []string{
				`File=C:\foo\boo`,
			},
			wantErr: false,
		},
		{
			name: "File= unix",
			fields: fields{
				connectionString: "File=/foo/boo",
			},
			want: []string{
				`File=/foo/boo`,
			},
			wantErr: false,
		},
		{
			name: "File=... quotes",
			fields: fields{
				connectionString: `File="C:\foo\boo"`,
			},
			want: []string{
				`File="C:\foo\boo"`,
			},
			wantErr: false,
		},
		{
			name: "/S have no sep",
			fields: fields{
				connectionString: "/Sfoo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "/S invalid sep",
			fields: fields{
				connectionString: "/Stcp://foo:1641/boo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "/S",
			fields: fields{
				connectionString: `/S127.0.0.1:1541\boo`,
			},
			want: []string{
				`Srvr=127.0.0.1:1541`,
				`Ref=boo`,
			},
			wantErr: false,
		},
		{
			name: "/S tcp",
			fields: fields{
				connectionString: `/Stcp://foo:1641\boo`,
			},
			want: []string{
				`Srvr=tcp://foo:1641`,
				`Ref=boo`,
			},
			wantErr: false,
		},
		{
			name: "/S IPv6",
			fields: fields{
				connectionString: `/S[fe10::c47b:90b7:fa32:a2fa%12]\boo`,
			},
			want: []string{
				`Srvr=[fe10::c47b:90b7:fa32:a2fa%12]`,
				`Ref=boo`,
			},
			wantErr: false,
		},
		{
			name: "/S multi-claster",
			fields: fields{
				connectionString: `/S127.0.0.1:1541,127.0.0.2:1542\boo`,
			},
			want: []string{
				`Srvr=127.0.0.1:1541,127.0.0.2:1542`,
				`Ref=boo`,
			},
			wantErr: false,
		},
		{
			name: "Srvr= Ref=",
			fields: fields{
				connectionString: `Srvr="foo";Ref=boo;`,
			},
			want: []string{
				`Srvr="foo"`,
				`Ref=boo`,
			},
			wantErr: false,
		},
		{
			name: "Ref= Srvr=",
			fields: fields{
				connectionString: `Ref=boo;Srvr="foo";`,
			},
			want: []string{
				`Ref=boo`,
				`Srvr="foo"`,
			},
			wantErr: false,
		},
		{
			name: "Srvr= Ref= tcp",
			fields: fields{
				connectionString: `Srvr=tcp://foo:1641;Ref=boo`,
			},
			want: []string{
				`Srvr=tcp://foo:1641`,
				`Ref=boo`,
			},
			wantErr: false,
		},
		{
			name: "Srvr= Ref= IPv6",
			fields: fields{
				connectionString: `Srvr=[fe10::c47b:90b7:fa32:a2fa%12];Ref=boo`,
			},
			want: []string{
				`Srvr=[fe10::c47b:90b7:fa32:a2fa%12]`,
				`Ref=boo`,
			},
			wantErr: false,
		},
		{
			name: "Srvr= Ref= multi-claster",
			fields: fields{
				connectionString: `Srvr=127.0.0.1:1541,127.0.0.2:1542;Ref=boo`,
			},
			want: []string{
				`Srvr=127.0.0.1:1541,127.0.0.2:1542`,
				`Ref=boo`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connectionString{
				connectionString: tt.fields.connectionString,
				values:           tt.fields.values,
			}
			err := c.parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(c.values, tt.want) {
				t.Errorf("parse() values = %v, want %v", c.values, tt.want)
			}
		})
	}
}
