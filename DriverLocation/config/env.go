package config

import "os"

const (
	envKey   = "APP_ENV"
	EnvLocal = "local_env"
	EnvTest = "test_env"
	EnvProd = "production_env"
)

var env = GetEnv(envKey, EnvLocal)

func Env() string {
	return env
}

func GetEnv(key, def string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}

	return def
}
