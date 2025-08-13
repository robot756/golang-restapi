package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"5"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() *Config {

	// Путь до конфиг-файла
	configPath := "./config/local.yaml"

	// Проверяем существование конфиг-файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exits: %s", configPath)
	}

	var cfg Config

	// Читаем конфиг-файл и заполняем нашу структуру
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return &cfg
}
