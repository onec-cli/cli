package platform

import (
	"reflect"
	"testing"
)

func Test_connStr_applayWithDefaults(t *testing.T) {
	type fields struct {
		connPath string
		values   []string
		connType connType
	}
	tests := []struct {
		name   string
		fields fields
		args   []string
		want   []string
	}{
		{
			name: "file connection",
			fields: fields{
				connPath: "/F./test",
				values:   nil,
				connType: File,
			},
			args: []string{"test"},
			want: nil,
		},
		{
			name: "client-server connection",
			fields: fields{
				connPath: `/Sfoo\boo`,
				values:   nil,
				connType: ClientServer,
			},
			args: []string{"param1=valueDefault1", "param2=valueDefault2"},
			want: []string{"param1=valueDefault1", "param2=valueDefault2"},
		},
		{
			name: "same parameters",
			fields: fields{
				connPath: `/Sfoo\boo`,
				values:   nil,
				connType: ClientServer,
			},
			args: []string{"param1=valueDefault1", "param1=valueDefault1"},
			want: []string{"param1=valueDefault1"},
		},
		{
			name: "new param",
			fields: fields{
				connPath: `/Sfoo\boo`,
				values:   []string{"param1=value1"},
				connType: ClientServer,
			},
			args: []string{"param2=valueDefault2"},
			want: []string{"param1=value1", "param2=valueDefault2"},
		},
		{
			name: "new doubled param",
			fields: fields{
				connPath: `/Sfoo\boo`,
				values:   []string{"param1=value1", "param2=value2", "param3=value3"},
				connType: ClientServer,
			},
			args: []string{"param2=valueDefault2"},
			want: []string{"param1=value1", "param2=value2", "param3=value3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connStr{
				connPath: tt.fields.connPath,
				values:   tt.fields.values,
				connType: tt.fields.connType,
			}
			c.apply(withDefaults(tt.args))
			if !reflect.DeepEqual(c.values, tt.want) {
				t.Errorf("defaultOptions() values = %v, want %v", c.values, tt.want)
			}
		})
	}
}

func Test_connStr_clean(t *testing.T) {
	type fields struct {
		values []string
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
			name: "have empty values",
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
			c := &connStr{
				values: tt.fields.values,
			}
			c.clean()
			if !reflect.DeepEqual(c.values, tt.want) {
				t.Errorf("clean() values = %v, want %v", c.values, tt.want)
			}
		})
	}
}

//func Test_makeFileStrings(t *testing.T) {
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name string
//		args args
//		want []string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := makeFileStrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("makeFileStrings() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_makeServerStrings(t *testing.T) {
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name string
//		args args
//		want []string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := makeServerStrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("makeServerStrings() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
