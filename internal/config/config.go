package config

/*
Package config
Подготавливает и загружает конфигурацию для служб сервиса.
*/

import (
	"fmt"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	CONFIG_DIR       = "configs"
	CONFIG_PROD_FILE = "prod"
	CONFIG_DEV_FILE  = "dev"
)

type Config struct {
	Environment Environment `mapstructure:"environment"`
	IsProd      bool
	IsDev       bool
	MariaDB     MariaDBConnections

	Http struct {
		Port int64 `mapstructure:"port"`
	} `mapstructure:"http"`

	Ctx struct {
		Ttl time.Duration `mapstructure:"ttl"`
	} `mapstructure:"ctx"`
}

type Environment struct {
	RootHost          string `mapstructure:"root_host"`
	RootRouter        string `mapstructure:"root_router"`
	PathToPublicFiles string `mapstructure:"path_to_public"`
}

type (
	MariaDBConnections struct {
		Default MariaDB
	}

	MariaDB struct {
		DB   string
		Host string
		Port int
		Name string
		User string
		Pass string
	}
)

func New(isProd string) (*Config, error) {
	// Check isProd
	if isProd == "" {
		return nil, fmt.Errorf("isProd is empty, please set it")
	}

	cfg := new(Config)

	// Set config path
	viper.AddConfigPath(CONFIG_DIR)
	viper.SetConfigName(CONFIG_DEV_FILE)

	// PRODUCTION
	if strings.Contains(isProd, CONFIG_PROD_FILE) {
		viper.SetConfigName(CONFIG_PROD_FILE)
		cfg.IsProd = true
	}

	// DEVELOP
	if strings.Contains(isProd, CONFIG_DEV_FILE) {
		viper.SetConfigName(CONFIG_PROD_FILE)
		cfg.IsDev = true
	}

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	// Load env vars
	if err := envconfig.Process("mariadb", &cfg.MariaDB.Default); err != nil {
		return nil, err
	}

	return cfg, nil
}
