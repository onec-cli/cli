package create

import (
	"errors"
	"github.com/v8platform/runner"
	"strings"
)

var errInvalidConnectionString = errors.New("invalid connection string format")

func NewInfobases(ib []string, opts ...string) []*infobase {
	var r []*infobase
	for _, c := range ib {
		command := &connectionString{connectionString: c}
		err := command.parse()
		if err == nil {
			command.defaultOptions(opts)
		}
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

type connType int

const (
	File connType = iota
	ClientServer
)

type connectionString struct {
	connectionString string
	values           []string
	connType         connType
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
	var values []string
	s := strings.Trim(c.connectionString, " ;")
	switch {
	case strings.HasPrefix(strings.ToUpper(s), "/F"):
		values = makeFileStrings(s)
	case strings.HasPrefix(strings.ToUpper(s), "/S"):
		c.connType = ClientServer
		values = makeServerStrings(s)
		if values == nil {
			return errInvalidConnectionString
		}
	case strings.Contains(strings.ToUpper(s), "FILE="):
		values = strings.Split(s, ";")
	case strings.Contains(strings.ToUpper(s), "SRVR="):
		c.connType = ClientServer
		values = strings.Split(s, ";")
	default:
		return errInvalidConnectionString
	}
	c.values = append(c.values, values...)
	c.removeEmpty()
	return nil
}

func (c *connectionString) defaultOptions(opts []string) {
	if c.connType != ClientServer {
		return
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
	srvr := "Srvr=" + strings.Trim(params[0], " ")
	ref := "Ref=" + strings.Trim(params[1], " ")
	return append(r, srvr, ref)
}
