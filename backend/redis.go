package backend

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	backend AbstractBackend
	client  *redis.Client

	Options *redis.Options
}

/**
 * Start
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *Redis) Start() {
	b.client = redis.NewClient(b.Options)
}

/**
 * Store a value
 * @return true if overwritting
 */
func (b *Redis) Store(key string, value string) error {
	err := b.client.HSet("links", key, value).Err()
	if err != nil {
		return err
	}
	err = b.client.HSet("metrics", key, 0).Err()
	return err
}

/**
 * Get ...
 * @return (url, true) if present
 *         (_, false) if no value present
 */
func (b *Redis) Get(key string) (string, error) {
	str, err := b.client.HGet("links", key).Result()
	return str, err
}

/**
 * Get ...
 * @return (url, true) if present
 *         (_, false) if no value present
 */
func (b *Redis) GetAll() (map[string]string, error) {
	str, err := b.client.HGetAll("links").Result()
	return str, err
}

/**
 * Delete ...
 * @return true if deleted
 *         false if not present
 */
func (b *Redis) Delete(key string) bool {
	_, err := b.client.HDel("links", key).Result()
	return err == nil
}

/**
 * Delete ...
 * @return true if deleted
 *         false if not present
 */
func (b *Redis) MetricIncrement(key string) {
	b.client.HIncrBy("metrics", key, 1)
}

/**
 * MetricGet
 */
func (b *Redis) MetricGet(key string) uint {
	if val, err := b.client.HGet("metrics", key).Uint64(); err == nil {
		return uint(val)
	}
	return 0
}
