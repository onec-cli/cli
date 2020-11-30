package create

import (
	"github.com/spf13/cast"
	"github.com/v8platform/marshaler"
	"reflect"
	"strings"
)

func DefaultOptions(opts map[string]interface{}) ([]string, error) {
	defOpts := newDefaultOptions()
	defOpts.bind(opts)
	return marshaler.Marshal(defOpts)
}

type defaultOptions struct {
	//тип используемого сервера баз данных:
	// MSSQLServer — Microsoft SQL Server;
	// PostgreSQL — PostgreSQL;
	// IBMDB2 — IBM DB2;
	// OracleDatabase — Oracle Database.
	DBMS string `v8:"DBMS, equal_sep" json:"dbms"`

	//имя сервера баз данных;
	DBSrvr string `v8:"DBSrvr, equal_sep" json:"db_srvr"`

	// имя базы данных в сервере баз данных;
	DB string `v8:"DB, equal_sep" json:"db_ref"`

	//имя пользователя сервера баз данных;
	DBUID string `v8:"DBUID, equal_sep" json:"db_user"`

	// создать базу данных в случае ее отсутствия ("Y"|"N".
	// "Y" — создавать базу данных в случае отсутствия,
	// "N" — не создавать. Значение по умолчанию — N).
	CrSQLDB bool `v8:"CrSQLDB, optional, equal_sep, bool_true=Y" json:"create_db"`
}

func newDefaultOptions() *defaultOptions {
	d := &defaultOptions{}
	return d
}

func (o *defaultOptions) bind(opts map[string]interface{}) {
	typeOf := reflect.TypeOf(*o)
	elem := reflect.ValueOf(o).Elem()
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		new := opts[strings.ToLower(field.Name)]
		target := elem.FieldByName(field.Name)
		switch target.Interface().(type) {
		case string:
			target.SetString(cast.ToString(new))
		case bool:
			target.SetBool(cast.ToBool(new))
		}
	}
}
