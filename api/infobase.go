package api

import (
	"errors"
	"github.com/v8platform/runner"
	"strings"
)

var ErrInvalidConnectionString = errors.New("invalid connection string format")

func CreateInfobase(c string) (runner.Command, error) {
	command := connectionString{connectionString: c}
	err := command.parse()
	if err != nil {
		return nil, err
	}
	return &command, nil
}

type connectionString struct {
	connectionString string
	values           []string
}

func (c *connectionString) Command() string {
	return runner.CreateInfobase
}

func (c *connectionString) Check() error {
	return nil
}

func (c *connectionString) Values() []string {
	return c.values
}

func (c *connectionString) removeEmpty() {
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

func (c *connectionString) parse() error {
	s := strings.Trim(c.connectionString, " ;")
	switch {
	case strings.HasPrefix(strings.ToUpper(s), "/F"):
		s = s[2:]
		s = "File=" + strings.Trim(s, " ")
		c.values = append(c.values, s)
	case strings.HasPrefix(strings.ToUpper(s), "/S"):
		s = s[2:]
		i := strings.LastIndex(s, "\\")
		if i < 0 {
			return ErrInvalidConnectionString
		}
		srv, ref := "Srvr="+strings.Trim(s[:i], " "), "Ref="+strings.Trim(s[i+1:], " ")
		c.values = append(c.values, srv, ref)
	case strings.Contains(strings.ToUpper(s), "FILE=") ||
		strings.Contains(strings.ToUpper(s), "SRVR="):
		c.values = strings.Split(s, ";")
		c.removeEmpty()
	default:
		return ErrInvalidConnectionString
	}
	return nil
}
