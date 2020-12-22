package cli

import (
	"github.com/onec-cli/cli/api"
	"io"
	"os"
)

// Cli represents the command line client.
type Cli interface {
	In() io.ReadCloser
	Out() io.Writer
	Err() io.Writer
	Platform() api.Platform
}

// cli is an instance the command line client.
// Instances of the client can be returned from NewCli.
type cli struct {
	in       io.ReadCloser
	out      io.Writer
	err      io.Writer
	platform api.Platform
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

// Platform returns the api.Platform
func (cli *cli) Platform() api.Platform {
	return cli.platform
}
