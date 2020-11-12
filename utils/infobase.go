package utils

import (
	"errors"
	"fmt"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
	"log"
	"reflect"
	"strings"
)

var _ v8.Command = (*CreateInfobase)(nil)

type CreateInfobase struct {
	ConnectString string
}

func (c CreateInfobase) Command() string {
	return runner.CreateInfobase
}

func (c CreateInfobase) Check() error {

	// TODO Сюда можно добавить любые проверки
	return nil
}

func (c CreateInfobase) Values() []string {

	var values []string

	switch {
	case strings.HasPrefix(strings.ToUpper(c.ConnectString), "/F"):

		values = append(values, "File="+strings.TrimLeft(c.ConnectString, "/F"))

	case strings.HasPrefix(strings.ToUpper(c.ConnectString), "/S"):

		path := strings.TrimLeft(c.ConnectString, "/S")
		r := strings.Replace(path, "\\", "/", 1)
		i := strings.LastIndex(r, "/")
		if i < 0 {
			log.Fatalf("invalid format for Srvr: %s", c.ConnectString)
		}
		values = append(values, "Srvr="+r[:i])
		values = append(values, "Ref="+r[i+1:])

		// TODO А это откуда брать в данном случае?
		//";DBMS=" + serverIB.DBMS +
		//";DBSrvr=" + serverIB.DBSrvr +
		//";DBUID=" + serverIB.DBUID +
		//";DBPwd=" + serverIB.DBPwd +
		//";DB=" + serverIB.DB +
		//";SQLYOffs=0"

	case strings.HasPrefix(c.ConnectString, "File=") ||
		strings.HasPrefix(c.ConnectString, "Srvr="):

		values = strings.Split(c.ConnectString, ";")
		// TODO Надо почистить от пустых строк и артифактов

	}

	return values
}

func (c CreateInfobase) Infobase() v8.Infobase {
	switch {
	case strings.HasPrefix(strings.ToUpper(c.ConnectString), "/F") ||
		strings.HasPrefix(c.ConnectString, "File="):

		var path string
		if strings.HasPrefix(strings.ToUpper(c.ConnectString), "/F") {
			path = strings.TrimLeft(c.ConnectString, "/F")
		} else {
			// чего то там сделать
		}

		return v8.FileInfoBase{
			//InfoBase: v8.InfoBase{},
			File: path,
			//Locale:   "", //TODO Добавить получение locale из строки
		}

	case strings.HasPrefix(strings.ToUpper(c.ConnectString), "/S") ||
		strings.HasPrefix(c.ConnectString, "Srvr="):

		var srvr, ref string

		if strings.HasPrefix(strings.ToUpper(c.ConnectString), "/S") {
			path := strings.TrimLeft(c.ConnectString, "/S")
			r := strings.Replace(path, "\\", "/", 1)
			i := strings.LastIndex(r, "/")
			if i < 0 {
				log.Fatalf("invalid format for Srvr: %s", c.ConnectString)
			}
			srvr, ref = r[:i], r[i+1:]
		} else {

			// чего то там сделать для строки подлчения вида "Srvr="
		}

		return v8.ServerInfoBase{
			Srvr: srvr,
			Ref:  ref,
		}
	}
	return nil
}

func NewInfobase(s string) (v8.Infobase, error) {

	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return nil, errors.New("the connection string is very short")
	}

	path := string([]rune(s)[2:])
	switch prefix := strings.ToUpper(s)[:2]; {
	case strings.HasPrefix(prefix, "/F"):
		ib := v8.NewFileIB(path)
		return &ib, nil
	case strings.HasPrefix(prefix, "/S"):
		srv, ref, err := getSrvrRef(path)
		if err != nil {
			return nil, fmt.Errorf("prefix %s: %v", prefix, err)
		}
		ib := v8.NewServerIB(srv, ref)
		return &ib, nil
	}

	if strings.Index(s, "File=") >= 0 || strings.Index(s, "Srvr=") >= 0 {
		params := splitParams(s)
		ib, err := getInfobase(params)
		if err != nil {
			return nil, fmt.Errorf("invalid connection string: %v", err)
		}
		err = fillInfobase(ib, params)
		if err != nil {
			return nil, fmt.Errorf("can't parse connection string: %v", err)
		}
		return ib, nil
	}
	return nil, errors.New("invalid connection string format")
}

func getInfobase(params map[string]interface{}) (v8.Infobase, error) {
	if _, ok := params["File"]; ok {
		return &v8.FileInfoBase{}, nil
	}
	if _, ok := params["Srvr"]; ok {
		return &v8.ServerInfoBase{}, nil
	}
	return nil, errors.New("must have param File or Srvr")
}

func splitParams(s string) map[string]interface{} {
	s = strings.TrimRight(s, ";")
	m := make(map[string]interface{})
	for _, params := range strings.Split(s, ";") {
		kv := strings.Split(params, "=")
		m[kv[0]] = strings.Trim(kv[1], "\"")
	}
	return m
}

func getSrvrRef(s string) (string, string, error) {
	r := strings.Replace(s, "\\", "/", 1)
	i := strings.LastIndex(r, "/")
	if i < 0 {
		return "", "", fmt.Errorf("invalid format for Srvr: %s", s)
	}
	return r[:i], r[i+1:], nil
}

func setField(obj interface{}, name string, value interface{}) error {

	elem := reflect.ValueOf(obj).Elem()
	fieldValue := elem.FieldByName(name)
	if !fieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}
	if !fieldValue.CanSet() {
		return fmt.Errorf("can't set %s field value", name)
	}
	fieldType := fieldValue.Type()
	v := reflect.ValueOf(value)
	if fieldType != v.Type() {
		return errors.New("provided value type didn't match obj field type")
	}
	fieldValue.Set(v)

	return nil
}

func fillInfobase(s v8.Infobase, m map[string]interface{}) error {
	for k, v := range m {
		err := setField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
