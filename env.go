package polarisb_config_go

import (
	"encoding/json"
	"expvar"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var env = expvar.NewMap("env")

// Get returns the value of the given environment variable.
func get(name string, defaultValue string) string {

	value, ok := os.LookupEnv(name)
	if !ok {
		value = defaultValue
	}

	env.Set(name, jsonStringer(value))

	return value
}

func set(name string, value string) {
	os.Setenv(name, value)
	env.Set(name, jsonStringer(value))
}

type jsonStringer string

func (s jsonStringer) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}

func tryLoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file \n ENV Load Error: %v \n", err)
	}
}
