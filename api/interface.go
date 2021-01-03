package api

// Infobase defines API methods for the infobase
type Infobase interface {
	DumpIB(file string) error
	Create() error
	Error() error
}
