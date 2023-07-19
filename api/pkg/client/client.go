package client

type Client interface {
	Post() error
	Get() error
}
