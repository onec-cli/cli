package build

const APP_NAME = "onec"

// Default build-time variable.
// These values are overridden via go build --ldflags
var (
	Version   = "unknown"
	GitCommit = "unknown"
	Time      = "unknown"
)
