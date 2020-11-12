package utils

import (
	"errors"
	"fmt"
	v8 "github.com/v8platform/v8"
	"reflect"
	"strings"
)

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
