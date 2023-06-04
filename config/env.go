package config

import (
	"os"
	"strings"
)

type Env struct {
	data map[string]string
}

func NewEnv() *Env {
	env := &Env{}

	env.data = make(map[string]string, 10)
	for _, pair := range os.Environ() {
		keyAndValue := strings.Split(pair, "=")
		env.data[keyAndValue[0]] = keyAndValue[1]
	}

	return env
}

func (e Env) GetValueByKey(key string) string {
	return e.data[key]
}
