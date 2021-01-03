package platform

import (
	"fmt"
	v8 "github.com/v8platform/api"
	"github.com/v8platform/runner"
	"log"
	"strings"
)

type infobase struct {
	connStr *connStr
	err     error
}

func (i *infobase) ConnectionString() string {
	connString := strings.Join(i.connStr.Values(), ";")
	return fmt.Sprintf("/IBConnectionString %s", connString)
	//panic("implement me")
}

func NewInfobase(connPath string, opts ...string) *infobase {
	c, err := NewConnStr(connPath)
	if err == nil {
		c.apply(withDefaults(opts))
	}
	return &infobase{
		connStr: c,
		err:     err,
	}
}

func (i *infobase) DumpIB(file string) error {
	if i.err != nil {
		return i.err
	}
	r := runner.NewPlatformRunner(i, v8.DumpIB(file))
	err := r.Run(nil)
	if err != nil {
		return err
	}
	return nil
}

func (i *infobase) Create() error {
	r := runner.NewPlatformRunner(nil, newCreateCommand(i.connStr.values))
	//	Spinner.Stop()
	log.Printf("=> %v\n", r.Args())
	//	Spinner.Start()
	err := r.Run(nil)
	if err != nil {
		return err
	}
	return nil
}

func (i *infobase) Error() error {
	return i.err
}

//func NewInfobases(ib []string, opts ...string) []*infobase {
//	var r []*infobase
//	for _, c := range ib {
//		connStr, err := NewConnStr(c)
//		if err == nil {
//			connStr.defaultOptions(opts)
//		}
//		r = append(r, newInfobase(connStr, err))
//	}
//	return r
//}

//func newInfobase(connStr *connStr, err error) *infobase {
//	return &infobase{connStr: connStr, err: err}
//}

type createCommand struct {
	values []string
}

func newCreateCommand(values []string) *createCommand {
	return &createCommand{values: values}
}

func (i *createCommand) Command() string {
	return runner.CreateInfobase
}

func (i *createCommand) Check() error {
	return nil
}

func (i *createCommand) Values() []string {
	return i.values
}

//
//func Test_connectionString_Command(t *testing.T) {
//	c := &connectionString{}
//	want := "CREATEINFOBASE"
//	got := c.Command()
//	if got != want {
//		t.Errorf("Command() = %v, want %v", got, want)
//	}
//}
//
//func Test_connectionString_Check(t *testing.T) {
//	c := &connectionString{}
//	got := c.Check()
//	if got != nil {
//		t.Errorf("Command() = %v, want %v", got, nil)
//	}
//}

//func (i *infobase) ConnectionString() string {
//	connString := strings.Join(i.connStr.Values(), ";")
//	return fmt.Sprintf("/IBConnectionString %s", connString)
//}
