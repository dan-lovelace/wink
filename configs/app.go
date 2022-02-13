package configs

type Config struct {
	DB  *DBConn
	Env *EnvConfig
}

type EnvConfig struct {
	Path string
	Type string
}

func GetAppConfig() Config {
	return Config{
		DB: &DBConn{
			Driver:   "sqlite3",
			Location: "./test.db",
		},
		Env: &EnvConfig{
			Path: "./.env",
			Type: "env",
		},
	}
}
