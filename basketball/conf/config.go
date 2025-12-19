package conf

import (
	"basketball/dao"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Manager struct {
	Port      int    `yaml:"port"`
	JWTSecret string `yaml:"jwtSecret"`
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
	dir, err := os.Getwd()
	fmt.Println("dir:", dir)
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName(commandArgs.ConfigFile)
	viper.AddConfigPath(filepath.Join(commandArgs.HomeDir, "/conf"))

	// Find and read the config file
	err = viper.ReadInConfig()

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
