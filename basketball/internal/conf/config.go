package conf

import (
	"basketball/internal/dao"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type Manager struct {
	Port          int    `yaml:"port"`
	JWTSecret     string `yaml:"jwtSecret"`
	Env           string `yaml:"env"`
	AccessExpire  int64  `yaml:"accessExpire"`
	RefreshExpire int64  `yaml:"refreshExpire"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Manager Manager `yaml:"manager"`
	Db      Db      `yaml:"db"`
	Redis   Redis   `yaml:"redis"`
}

func InitConfig(commandArgs *dao.CommandArgs) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigName(commandArgs.ConfigFile)
	viper.AddConfigPath(filepath.Join(commandArgs.HomeDir, "/internal/conf"))

	// Find and read the config file
	err := viper.ReadInConfig()

	// Handle errors
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal config file: %w", err))
	}

	return config
}
