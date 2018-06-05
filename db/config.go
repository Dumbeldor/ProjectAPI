package db

// UsersDBConfig configuration object to use in YAML configurations
type UsersDBConfig struct {
	URL          string `yaml:"url"`
	MaxIdleConns int    `yaml:"max-idle-conns"`
	MaxOpenConns int    `yaml:"max-open-conns"`
}
