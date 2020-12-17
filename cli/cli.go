package cli

import (
	"io"
	"os"
)

// Cli represents the command line client.
type Cli interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer
}

// cli is an instance the command line client.
// Instances of the client can be returned from NewCli.
type cli struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

// NewCli returns a cli instance with all operators applied on it.
// It applies by default the standard streams.
func NewCli() *cli {
	cli := &cli{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
	return cli
}

// In returns the reader used for stdin
func (cli *cli) In() io.Reader {
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
