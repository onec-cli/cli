package platform

import (
	"fmt"
	v8 "github.com/v8platform/api"
	"github.com/v8platform/runner"
)

type infobase struct {
	connStr *connStr
	err     error
}

func (i *infobase) ConnectionString() string {
	panic("implement me")
}

func NewInfobase(connPath string, opts ...string) *infobase {
	c, err := NewConnStr(connPath)
	if err == nil {
		c.defaultOptions(opts)
	}
	return &infobase{
		connStr: c,
		err:     err,
	}
}

func (i *infobase) DumpIB(file string) error {
	r := runner.NewPlatformRunner(i, v8.DumpIB(file))
	err := r.Run(nil)
	if err != nil {
		return err
	}
	return nil
}

func (i *infobase) Create() {
	fmt.Println(i.Values())
}

func (i *infobase) Error() error {
	return i.err
}

func NewInfobases(ib []string, opts ...string) []*infobase {
	var r []*infobase
	for _, c := range ib {
		connStr, err := NewConnStr(c)
		if err == nil {
			connStr.defaultOptions(opts)
		}
		r = append(r, newInfobase(connStr, err))
	}
	return r
}

func newInfobase(connStr *connStr, err error) *infobase {
	return &infobase{connStr: connStr, err: err}
}

func (i *infobase) Command() string {
	return runner.CreateInfobase
}

func (i *infobase) Check() error {
	return nil
}

func (i *infobase) Values() []string {
	return i.connStr.values
}

//func (i *infobase) ConnectionString() string {
//	connString := strings.Join(i.connStr.Values(), ";")
//	return fmt.Sprintf("/IBConnectionString %s", connString)
//}

func (i *infobase) Err() error {
	return i.err
}
