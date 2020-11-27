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

func (c *connectionString) parse() error {
	s := strings.Trim(c.connectionString, " ;")
	switch {
	case strings.HasPrefix(strings.ToUpper(s), "/F"):
		s = s[2:]
		s = "File=" + strings.Trim(s, " ")
		c.values = append(c.values, s)
	case strings.HasPrefix(strings.ToUpper(s), "/S"):
		s = s[2:]
		if i := strings.LastIndex(s, "\\"); i > 0 {
			c.values = append(c.values, "Srvr="+strings.Trim(s[:i], " "), "Ref="+strings.Trim(s[i+1:], " "))
		} else {
			return ErrInvalidConnectionString
		}
	case strings.Contains(s, "File=") || strings.Contains(s, "Srvr="):
		c.values = strings.Split(s, ";")
		b := c.values[:0]
		for _, x := range c.values {
			if x != "" {
				b = append(b, x)
			}
		}
	default:
		return ErrInvalidConnectionString
	}
	return nil
}
