package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Структуры, соответствующие конфигурации ../../config/local.yaml
type Config struct {
	Env         string `yaml:"env" env-default: "local"`         // установка значений по умолчанию
	StoragePath string `yaml:"storage_path" env-required:"true"` // тег env-required, чтобы приложение не запустилось в случае отсутсвия пути storage
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// Чтение файла конфигурации ../../config/local.yaml, создание и заполнение объекта Config (инициализация)
// Приставка Must в случае возврата ошибки вызывает panic()
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
