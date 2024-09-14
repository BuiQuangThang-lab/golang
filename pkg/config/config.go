package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"sync"
)

type App struct {
	Mysql Mysql `mapstructure:"mysql"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

var (
	configSync    sync.Once
	currentConfig *App
)

func LoadConfig(filepath ...string) *App {
	configSync.Do(func() {
		if len(filepath) == 0 {
			fmt.Println("No configuration file provided")
			return
		}

		configFile := filepath[0]
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			fmt.Printf("Configuration file not found: %s\n", configFile)
			return
		}

		v := viper.New()
		v.SetConfigFile(configFile)
		v.SetConfigType("yaml")

		// Đọc file cấu hình
		err := v.ReadInConfig()
		if err != nil {
			fmt.Printf("Error reading config file: %v\n", err)
			return
		}

		// Giải mã cấu hình vào struct App
		currentConfig = &App{}
		err = v.UnmarshalExact(currentConfig)
		if err != nil {
			fmt.Printf("Error unmarshalling config: %v\n", err)
			currentConfig = nil
			return
		}
	})

	if currentConfig == nil {
		fmt.Println("Failed to load configuration")
	}

	return currentConfig
}
