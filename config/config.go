package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	URI string `yaml:"uri"`
}

type LoggingConfig struct {
	Level string `yaml:"level"`
}

type Config struct {
	Server       ServerConfig   `yaml:"server"`
	Database     DatabaseConfig `yaml:"database"`
	Logging      LoggingConfig  `yaml:"logging"`
	JwtSecretKey string         `yaml:"jwt_secret_key"`
	JwtIssuer    string         `yaml:"jwt_issuer"`
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig(configFilePath string) error {
	var err error
	once.Do(func() {
		file, e := os.Open(configFilePath)
		if e != nil {
			err = fmt.Errorf("could not open config file: %v", e)
			return
		}
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		config := &Config{}
		if e := decoder.Decode(config); e != nil {
			err = fmt.Errorf("could not decode config file: %v", e)
			return
		}

		cfg = config
	})
	return err
}

func GetConfig() *Config {
	return cfg
}
