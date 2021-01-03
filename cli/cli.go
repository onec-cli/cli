package cli

import (
	"github.com/onec-cli/cli/api"
	"github.com/onec-cli/cli/internal/platform"
	"io"
	"os"
)

// Cli represents the command line client.
type Cli interface {
	In() io.ReadCloser
	Out() io.Writer
	Err() io.Writer
	Infobase(connPath string, opts ...string) api.Infobase
}

// cli is an instance the command line client.
// Instances of the client can be returned from NewCli.
type cli struct {
	in  io.ReadCloser
	out io.Writer
	err io.Writer
	ib  map[string]api.Infobase
}

// NewCli returns a cli instance with all operators applied on it.
// It applies by default the standard streams.
func NewCli() *cli {
	cli := &cli{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
		ib:  make(map[string]api.Infobase),
	}
	return cli
}

// In returns the reader used for stdin
func (cli *cli) In() io.ReadCloser {
	return cli.in
}

// Out returns the writer used for stdout
func (cli *cli) Out() io.Writer {
	return cli.out
}

// Err returns the writer used for stderr
func (cli *cli) Err() io.Writer {
	return cli.err
}

//todo
func (cli *cli) Infobase(connPath string, opts ...string) api.Infobase {
	if ib, ok := cli.ib[connPath]; ok {
		return ib
	}
	cli.ib[connPath] = platform.NewInfobase(connPath, opts...)
	return cli.ib[connPath]
}
