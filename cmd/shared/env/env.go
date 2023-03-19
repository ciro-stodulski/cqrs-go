package env

import (
	"github.com/thrcorrea/envloader"
)

type Environment struct {
	HostPort string `env:"HOST_PORT"`
	HostHttp string `env:"HOST_HTTP"`
}

var env Environment

func Env() *Environment {
	return &env
}

func Load() {
	err := envloader.Load(&env, ".env")
	if err != nil {
		panic(err)
	}
}
