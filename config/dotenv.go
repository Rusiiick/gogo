package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const configPath = "../config.yaml"

type DBConfig struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"Dbname"`
}

type Config struct {
	JWTSecret            string        `yaml:"JWTSecret"`
	DBconfig             DBConfig      `yaml:"DBConfig"`
	AccessTokenDuration  time.Duration `yaml:"AccessTokenDuration"`
	RefreshTokenDuration time.Duration `yaml:"RefreshTokenDuration"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %v", err)
		return nil, err
	}

	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Fatalf("cannot read config: %v", err)
		return nil, err
	}

	return config, nil
}
