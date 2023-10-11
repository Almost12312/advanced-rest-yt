package config

import (
	"advanced-rest-yt/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Env     string `yaml:"env" env-required:"true"`
	Listen  listen `yaml:"listen" env-required:"true"`
	MongoDB MongoDB
}

type listen struct {
	Type   string `yaml:"type"`
	BindIp string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type MongoDB struct {
	Port       string `json:"port"`
	Collection string `json:"collection"`
	Database   string `json:"database"`
	AuthDB     string `json:"auth_db"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("start read config!")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatalln("Cant read config info", err)
		}
	})

	return instance
}
