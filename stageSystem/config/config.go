package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

/***
* usage example:
* config.GoDotEnvVariable("STRONGEST_AVENGER")
 */

const (
	ENV_FILE_NAME = ".env"
)

func GoDotEnvVariable(key string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("caller error")
	}
	envFile := fmt.Sprintf("%s/%s", path.Dir(filename), ENV_FILE_NAME)

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("reading config file error (.env file) key:", key)
	}

	return os.Getenv(key)
}