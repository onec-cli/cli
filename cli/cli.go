package cli

import (
	"io"
	"os"
)

// Cli is an instance the command line client.
// Instances of the client can be returned from NewCli.
type Cli struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

// NewCli returns a cli instance with all operators applied on it.
// It applies by default the standard streams.
func NewCli() *Cli {
	cli := &Cli{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
	return cli
}

// Out returns the writer used for stdout
func (cli *Cli) Out() io.Writer {
	return cli.out
}

// Err returns the writer used for stderr
func (cli *Cli) Err() io.Writer {
	return cli.err
}

// In returns the reader used for stdin
func (cli *Cli) In() io.Reader {
	return cli.in
}
