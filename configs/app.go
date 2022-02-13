package configs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	CurrentProject string `mapstructure:"CURRENT_PROJECT"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBLocation     string `mapstructure:"DB_LOCATION"`
}

func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cfgName := "config"
	cfgType := "env"
	cfgPath := path.Join(homeDir, ".wink", fmt.Sprintf("%s.%s", cfgName, cfgType))
	cfgDir := filepath.Dir(cfgPath)

	// create an empty config file if it doesn't exist
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		if err := os.MkdirAll(cfgDir, 0770); err != nil {
			return nil, err
		}

		os.Create(cfgPath)
	}

	// set up viper config
	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(cfgDir)

	// viper defaults
	viper.SetDefault("CURRENT_PROJECT", "")
	viper.SetDefault("DB_DRIVER", "sqlite3")
	viper.SetDefault("DB_LOCATION", path.Join(cfgDir, "wink_production.db"))

	// load config for output
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// write config to file
	if err := viper.WriteConfig(); err != nil {
		return nil, err
	}

	// construct output
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
