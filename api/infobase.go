package api

import (
	"errors"
	"github.com/v8platform/runner"
	"strings"
)

var errInvalidConnectionString = errors.New("invalid connection string format")

func CreateInfobase(s []string) []*infobase {
	var r []*infobase
	for _, c := range s {
		command := &connectionString{connectionString: c}
		err := command.parse()
		r = append(r, newInfobase(command, err))
	}
	return r
}

type infobase struct {
	command runner.Command
	err     error
}

func newInfobase(command runner.Command, err error) *infobase {
	return &infobase{command: command, err: err}
}

func (i *infobase) Command() (runner.Command, error) {
	return i.command, i.err
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
			return errInvalidConnectionString
		}
		srv, ref := "Srvr="+strings.Trim(s[:i], " "), "Ref="+strings.Trim(s[i+1:], " ")
		c.values = append(c.values, srv, ref)
	case strings.Contains(strings.ToUpper(s), "FILE=") ||
		strings.Contains(strings.ToUpper(s), "SRVR="):
		c.values = strings.Split(s, ";")
		c.removeEmpty()
	default:
		return errInvalidConnectionString
	}
	return nil
}
