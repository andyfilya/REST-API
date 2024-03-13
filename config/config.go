package config

import (
	"os"
  "errors"

	"gopkg.in/yaml.v2"
)

const (
  configDirectory = "./config"
  configYaml = "config.yaml"
)


type ServerConfig struct {
  Host string `yaml:"host"`
  Port string `yaml:"port"`
}

type GlobalConfig struct {
  ServCfg ServerConfig `yaml:"server"` 
}

func InitGlobalConfig() (*GlobalConfig, error) {
  gc := &GlobalConfig{}

  bytes, err := os.ReadFile(configDirectory+"/"+configYaml)
  
  if err != nil {
    return nil, errors.New("error while reading directory")
  }

  err = yaml.Unmarshal(bytes, &gc)
  if err != nil {
    return nil, errors.New("error while unmarshal yaml file")
  }

  return gc, nil
}
