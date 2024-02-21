package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type NatsStreamingConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	ClusterID string `yaml:"сlusterID"`
	ClientID  string `yaml:"clientID"`
}

type AppConfig struct {
	Database StorageConfig       `yaml:"database"`
	Stan     NatsStreamingConfig `yaml:"stan"`
}

func New() (*AppConfig, error) {
	var cfg AppConfig

	file, err := os.Open(".././config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
