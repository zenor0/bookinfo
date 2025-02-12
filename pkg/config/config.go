package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Tracing  TracingConfig  `mapstructure:"tracing"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type TracingConfig struct {
	Enabled    bool   `mapstructure:"enabled"`
	JaegerURL  string `mapstructure:"jaeger_url"`
	ServiceName string `mapstructure:"service_name"`
}

func LoadConfig(configPath string) (*Config, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Failed to read config file", zap.Error(err))
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.Error("Failed to unmarshal config", zap.Error(err))
		return nil, err
	}

	return &config, nil
} 