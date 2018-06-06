package internal

import (
	"gitlab.com/projetAPI/auth"
	"gitlab.com/projetAPI/ProjetAPI/service"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"gitlab.com/projetAPI/ProjetAPI/db"
)

type config struct {
	UsersDB db.UsersDBConfig `yaml:"usersdb"`
	Redis auth.RedisConfig
	HTTP service.HTTPConfig `yaml:"http"`
	Log service.LogConfig `yaml:"log"`
}

var gconfig config

func (c *config) loadDefaultConfiguration() {
	c.UsersDB.URL = "host=postgres dbname=postgres user=postgres password=mysecretpassword"
	c.UsersDB.MaxOpenConns = 10
	c.UsersDB.MaxIdleConns = 5

	c.Redis.Host = "redis"
	c.Redis.Port = 6379
	c.Redis.Password = ""
	c.Redis.DatabaseID = 0
	c.Redis.MaxRetries = 3

	c.HTTP.Port = 8080
	c.HTTP.EnableCORS = false

	c.Log.EnableSyslog = false
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
		app.Log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		app.Log.Fatalf("error: %v", err)
	}

	println(fmt.Sprintf("Configuration loaded from '%s'.", path))
	return true
}
