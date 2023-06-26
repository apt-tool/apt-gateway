package redis

type Connector interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
