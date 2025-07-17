package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port       string   `mapstructure:"port"`
		AuthTokens []string `mapstructure:"auth_tokens"`
	} `mapstructure:"server"`

	API struct {
		BaseURL string `mapstructure:"base_url"`
	} `mapstructure:"api"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")   // имя конфиг файла (без расширения)
	viper.SetConfigType("yaml")     // формат конфига
	viper.AddConfigPath(".")        // путь к текущей директории
	viper.AddConfigPath("./config") // путь к папке config

	// Чтение конфига
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Установка значений по умолчанию
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}

	return &cfg, nil
}
