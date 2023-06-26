package redis

import "time"

type Connector interface {
	Get(key string) (string, error)
	Set(key, value string, expire time.Duration) error
}
