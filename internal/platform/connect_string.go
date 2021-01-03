package platform

import (
	"errors"
	"strings"
)

var errInvalidConnStr = errors.New("connection string: invalid format")

// connStrOption applies a modification on a connStr.
type connStrOption func(c *connStr) error

type connType int

const (
	File connType = iota
	ClientServer
)

type connStr struct {
	connPath string
	connType connType
	values   []string
}

func NewConnStr(connPath string) (*connStr, error) {
	c := &connStr{connPath: connPath}
	if err := c.parse(); err != nil {
		return nil, err
	}
	c.clean()
	return c, nil
}

func (c *connStr) Values() []string {
	return c.values
}

func (c *connStr) parse() error {
	var values []string
	s := strings.Trim(c.connPath, " ;")
	switch {
	case strings.HasPrefix(strings.ToUpper(s), "/F"):
		values = makeFileStrings(s)
	case strings.HasPrefix(strings.ToUpper(s), "/S"):
		c.connType = ClientServer
		values = makeServerStrings(s)
		if values == nil {
			return errInvalidConnStr
		}
	case strings.Contains(strings.ToUpper(s), "FILE="):
		values = strings.Split(s, ";")
	case strings.Contains(strings.ToUpper(s), "SRVR="):
		c.connType = ClientServer
		values = strings.Split(s, ";")
	default:
		return errInvalidConnStr
	}
	c.values = append(c.values, values...)
	return nil
}

func (c *connStr) clean() {
	if c.values == nil {
		return
	}
	n := c.values[:0]
	for _, v := range c.values {
		if v != "" {
			n = append(n, v)
		}
	}
	c.values = n
}

func makeFileStrings(s string) []string {
	var r []string
	s = s[2:]
	return append(r, "File="+strings.Trim(s, " "))
}

func makeServerStrings(s string) []string {
	var r []string
	s = s[2:]
	params := strings.Split(s, "\\")
	if len(params) != 2 {
		return nil
	}
	srvr := "Srvr=" + strings.TrimSpace(params[0])
	ref := "Ref=" + strings.TrimSpace(params[1])
	return append(r, srvr, ref)
}

// apply all the operation on the connStr
func (c *connStr) apply(opts ...connStrOption) error {
	for _, op := range opts {
		if err := op(c); err != nil {
			return err
		}
	}
	return nil
}

// withDefaults appends a fragment <param>=<value> to the connection string values, skipping the addition
// if the fragment already exists. Only for client-server connection.
func withDefaults(opts []string) connStrOption {
	return func(c *connStr) error {
		if c.connType != ClientServer {
			return nil
		}
	exit:
		for _, s := range opts {
			params := strings.SplitAfter(s, "=")
			for _, value := range c.values {
				if strings.HasPrefix(value, params[0]) {
					continue exit
				}
			}
			c.values = append(c.values, s)
		}
		return nil
	}
}
