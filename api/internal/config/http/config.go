package http

type Config struct {
	Port int    `koanf:"port"`
	Core string `koanf:"core"`
}
