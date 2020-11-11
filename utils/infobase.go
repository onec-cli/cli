package utils

import (
	v8 "github.com/v8platform/v8"
	"strings"
)

func GetInfobase(s string) v8.Infobase {

	s = strings.TrimSpace(s)

	if len(s) < 2 {
		return nil
	}

	switch prefix := strings.ToUpper(s)[:2]; {
	case strings.HasPrefix(prefix, "/F"):
		return v8.NewFileIB(string([]byte(s)[2:]))
	case strings.HasPrefix(prefix, "/S"):
		s = strings.Replace(s, "\\", "/", 1)
		p := strings.Split(string([]byte(s)[2:]), "/")
		if len(p) != 2 {
			return nil
		}
		return v8.NewServerIB(p[0], p[1])
	default:
		return nil
	}

	//p := strings.Split(s, ";")
	//m := make(map[string]string)
	//for _, pair := range p {
	//	z := strings.Split(pair, "=")
	//	m[z[0]] = z[1]
	//}

}
