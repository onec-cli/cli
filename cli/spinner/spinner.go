package spinner

import (
	"github.com/briandowns/spinner"
	"time"
)

var Spinner *spinner.Spinner

func init() {
	Spinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
}
