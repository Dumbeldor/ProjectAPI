package service

// HTTPConfig standard http configuration object
type HTTPConfig struct {
	Port       int16 `yaml:"port"`
	EnableCORS bool  `yaml:"enable-cors"`
}

// LogConfig standard log configuration object
type LogConfig struct {
	EnableSyslog bool   `yaml:"enable-syslog"`
	LogLevel     string `yaml:"loglevel"`
}
