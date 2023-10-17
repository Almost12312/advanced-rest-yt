package config

import (
	"advanced-rest-yt/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Env     string  `yaml:"env" env-required:"true"`
	Listen  listen  `yaml:"listen" env-required:"true"`
	Storage Storage `yaml:"storage"`
}

type listen struct {
	Type   string `yaml:"type"`
	BindIp string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type Storage struct {
	MongoDB    MongoDB    `json:"mongodb" yaml:"mongodb"`
	PostgreSQL PostgreSQL `json:"postgresql" yaml:"postgres"`
}

type PostgreSQL struct {
	PreHost  string `json:"prehost" yaml:"pre_host"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
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
