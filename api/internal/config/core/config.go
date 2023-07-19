package core

type Config struct {
	Timeout    int  `koanf:"timeout"`
	Port       int  `koanf:"port"`
	Enable     bool `koanf:"enable"`
	Preemptive bool `koanf:"preemptive"`
}
