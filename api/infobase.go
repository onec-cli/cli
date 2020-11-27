package api

import (
	"errors"
	"github.com/v8platform/runner"
	"strings"
)

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

	switch {
	case strings.HasPrefix(strings.ToUpper(c.connectionString), "/F"):
		file := "File=" + strings.TrimPrefix(c.connectionString, "/F")
		c.values = append(c.values, file)
	case strings.Contains(c.connectionString, "File=") ||
		strings.Contains(c.connectionString, "Srvr="):
		c.values = strings.Split(c.connectionString, ";")
		// TODO Надо почистить от пустых строк и артифактов
	default:
		return errors.New("invalid connection string format")
	}

	return nil
}
