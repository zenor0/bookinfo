package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Tracing  TracingConfig  `mapstructure:"tracing"`
	Services ServicesConfig `mapstructure:"services"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
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
	Enabled     bool   `mapstructure:"enabled"`
	JaegerURL   string `mapstructure:"jaeger_url"`
	ServiceName string `mapstructure:"service_name"`
}

type ServicesConfig struct {
	Details ServiceEndpoint `mapstructure:"details"`
	Reviews ServiceEndpoint `mapstructure:"reviews"`
	Ratings ServiceEndpoint `mapstructure:"ratings"`
}

type ServiceEndpoint struct {
	URL string `mapstructure:"url"`
}

func LoadConfig(configDir string) (*Config, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 获取环境，默认为test
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "test"
	}

	// 构建配置文件路径
	configPath := filepath.Join(configDir, fmt.Sprintf("%s.yaml", env))

	// 设置基本配置
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 读取配置
	if err := v.ReadInConfig(); err != nil {
		logger.Error("Failed to read config file",
			zap.String("path", configPath),
			zap.Error(err))
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		logger.Error("Failed to unmarshal config", zap.Error(err))
		return nil, err
	}

	// 允许通过环境变量覆盖数据库配置
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.Database.User = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.Database.Password = dbPassword
	}

	return &config, nil
}
