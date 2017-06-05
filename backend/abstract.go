package backend

/**
 * AbstractBackend is the interface of all backends
 */
type AbstractBackend interface {
	Start()
	Store(string, string) error
	Get(string) (string, error)
	Delete(string) bool
	MetricIncrement(string)
	MetricGet(string) uint
}

var ActiveBackend AbstractBackend
