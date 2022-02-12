package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *DBConn
}

type Env interface {
	Get() string
	Set(value string)
}

func GetAppConfig() Config {
	return Config{
		DB: &DBConn{
			Driver:   "sqlite3",
			Location: "./test.db",
		},
	}
}

func GetEnv(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	val := os.Getenv(key)

	return val, nil
}

func SetEnv(key string, value string) error {
	env, err := godotenv.Unmarshal(fmt.Sprintf("%s=%s", key, value))
	if err != nil {
		return err
	}

	wErr := godotenv.Write(env, "./.env")
	if wErr != nil {
		return err
	}

	return nil
}
