package configs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	CurrentProject = "CURRENT_PROJECT"
	DBDriver       = "DB_DRIVER"
	DBLocation     = "DB_LOCATION"
	InitDir        = "INIT_DIR"
	TestDBLocation = "TEST_DB_LOCATION"
)

type Config struct {
	CurrentProject string `mapstructure:"CURRENT_PROJECT"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBLocation     string `mapstructure:"DB_LOCATION"`
	InitDir        string `mapstructure:"INIT_DIR"`
	TestDBLocation string `mapstructure:"TEST_DB_LOCATION"`
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
	dbDriver := "sqlite3"
	viper.SetDefault(CurrentProject, "")
	viper.SetDefault(DBDriver, dbDriver)
	viper.SetDefault(DBLocation, path.Join(cfgDir, fmt.Sprintf("wink_production.%s", dbDriver)))
	viper.SetDefault(TestDBLocation, path.Join(cfgDir, fmt.Sprintf("wink_test.%s", dbDriver)))
	viper.SetDefault(InitDir, "")

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
