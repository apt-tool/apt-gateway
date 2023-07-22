package http

type Config struct {
	Port    int    `koanf:"port"`
	Core    string `koanf:"core"`
	FTPHost string `koanf:"ftp_host"`
	DevMode bool   `koanf:"dev_mode"`
}
