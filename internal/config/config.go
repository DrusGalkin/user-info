package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Env   string       `yaml:"env" env-default:"local"`
	Sever ServerConfig `yaml:"server"`
	GRPC  GRPCConfig   `yaml:"grpc"`
	Redis RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
	Port    int           `yaml:"port" env-default:"8080"`
}

type GRPCConfig struct {
	Timeout time.Duration `yaml:"timeout"`
	Port    int           `yaml:"port"`
}

type RedisConfig struct {
	Address string        `yaml:"address" env-default:"localhost"`
	Port    int           `yaml:"port" env-default:"6379"`
	TTL     time.Duration `yaml:"ttl" env-default:"5m"`
}

func MustLoadConfig() *Config {
	const op = "config.MustLoadConfig"

	var cfg Config
	path := fetchConfigPath()

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	return &cfg
}

func fetchConfigPath() string {
	const op = "config.MustLoadEnv"
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	return os.Getenv("CONFIG_PATH")
}
