package http

type Config struct {
	Port       int    `koanf:"port"`
	Core       string `koanf:"core"`
	CoreSecret string `koanf:"core_secret"`
	DevMode    bool   `koanf:"dev_mode"`
}
