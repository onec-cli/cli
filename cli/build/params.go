package build

const AppName = "onec" // todo заменяются ли константы --ldflags?

// Default build-time variable.
// These values are overridden via go build --ldflags
var (
	Version = "unknown"
	Commit  = "unknown"
	Time    = "unknown"
)
