package backend

import (
	"errors"
	"sync"
)

type Dict struct {
	backend AbstractBackend
	links   map[string]string
	metrics map[string]uint

	linksLock   sync.RWMutex
	metricsLock sync.Mutex
}

/**
 * Start
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *Dict) Start() {
	b.links = make(map[string]string)
	b.metrics = make(map[string]uint)
	b.linksLock = sync.RWMutex{}
	b.metricsLock = sync.Mutex{}
}

/**
 * Store a value
 * @return true if overwritting
 */
func (b *Dict) Store(key string, value string) error {
	b.linksLock.Lock()
	defer b.linksLock.Unlock()
	b.links[key] = value
	b.metrics[key] = 0
	return nil
}

/**
 * Get ...
 * @return (url, true) if present
 *         (_, false) if no value present
 */
func (b *Dict) Get(key string) (string, error) {
	b.linksLock.RLock()
	defer b.linksLock.RUnlock()
	if str, ok := b.links[key]; ok {
		return str, nil
	} else {
		return "", errors.New("No key found.")
	}
}

/**
 * Delete ...
 * @return true if deleted
 *         false if not present
 */
func (b *Dict) Delete(key string) bool {
	b.linksLock.Lock()
	defer b.linksLock.Unlock()
	_, ok := b.links[key]
	if ok {
		delete(b.links, key)
	}
	return ok
}

/**
 * Delete ...
 * @return true if deleted
 *         false if not present
 */
func (b *Dict) MetricIncrement(key string) {
	b.metricsLock.Lock()
	defer b.metricsLock.Unlock()
	if _, ok := b.links[key]; ok {
		b.metrics[key]++
	}
}

/**
 * MetricGet
 */
func (b *Dict) MetricGet(key string) uint {
	b.metricsLock.Lock()
	defer b.metricsLock.Unlock()
	if val, ok := b.metrics[key]; ok {
		return val
	} else {
		return 0
	}
}
