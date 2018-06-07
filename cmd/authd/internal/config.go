package internal

import (
	"fmt"
	"gitlab.com/projetAPI/ProjetAPI/service"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	Postgresql struct {
		Connstr string `yaml:"connstr"`
	}
	Redis   service.RedisConfig
	HTTP    service.HTTPConfig `yaml:"http"`
	Log     service.LogConfig  `yaml:"log"`
	Session struct {
		Duration int64 `yaml:"duration"`
	}
}

var gconfig config

func (c *config) loadDefaultConfiguration() {
	c.Postgresql.Connstr = "host=127.0.0.1 dbname=postgres user=postgres password=example sslmode=disable"

	c.Redis.Host = "127.0.0.1"
	c.Redis.Port = 6379
	c.Redis.Password = ""
	c.Redis.DatabaseID = 0
	c.Redis.MaxRetries = 3

	c.HTTP.Port = 8080
	c.HTTP.EnableCORS = false

	c.Log.EnableSyslog = false

	c.Session.Duration = 86400
}

func (c *config) load(path string) bool {
	c.loadDefaultConfiguration()

	if len(path) == 0 {
		println("Configuration path is empty, using default configuration.")
		return false
	}

	println(fmt.Sprintf("Loading configuration from '%s'...", path))

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		println(fmt.Sprintf("Failed to read YAML file #%v", err))
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	println(fmt.Sprintf("Configuration loaded from '%s'.", path))
	return true
}
