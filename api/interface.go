package api

// Platform defines API methods for the 1C:Enterprise platform
type Platform interface {
	//DumpIB(ctx context.Context, file string) (*os.File, error)
	DumpIB(file string) error
}
