package config

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configYaml = "../../config/config.yaml"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type UserDatabaseConfig struct {
	DatabaseName string `yaml:"dbname"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"user"`
	Password     string `yaml:"password"`
	SSLmode      string `yaml:"sslmode"`
}

type GlobalConfig struct {
	ServCfg         ServerConfig       `yaml:"server"`
	UserDatabaseCfg UserDatabaseConfig `yaml:"database"`
}

func InitGlobalConfig() (*GlobalConfig, error) {
	gc := &GlobalConfig{}

	bytes, err := os.ReadFile(configYaml)

	if err != nil {
		logrus.Errorf("error reading directory : [%v]", err)
		return nil, errors.New("error while reading directory")
	}

	err = yaml.Unmarshal(bytes, &gc)
	if err != nil {
		logrus.Errorf("error unmarshal yaml config : [%v]", err)
		return nil, errors.New("error while unmarshal yaml file")
	}

	return gc, nil
}
