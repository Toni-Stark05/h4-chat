package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/samber/do"
	"log"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env"`
	HTTPServer `yaml:"http_server"`
	Postgres   `yaml:"postgres"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func mustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannon read config: %s", err)
	}

	return &cfg
}

func NewConfig(i *do.Injector) (*Config, error) {
	logger := do.MustInvoke[*slog.Logger](i)
	cfg := mustLoad()
	logger.Info("Config initialized", slog.String("env", cfg.Env))
	return cfg, nil
}
