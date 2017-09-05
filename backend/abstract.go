package backend

/**
 * AbstractBackend is the interface of all backends
 */
type AbstractBackend interface {
	Start()
	GetAll() (map[string]string, error)
	Store(string, string) error
	Get(string) (string, error)
	Delete(string) bool
	MetricIncrement(string)
	MetricGet(string) uint
}

var ActiveBackend AbstractBackend
