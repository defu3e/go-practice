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
* Пример использования:
* config.GoDotEnvVariable("HOST")
 */

const (
	ENV_FILE_NAME = ".env"
)

func GoDotEnvVariable(key string) string {
	// получаем путь к env-файлу
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("caller error")
	}
	envFile := fmt.Sprintf("%s/%s", path.Dir(filename), ENV_FILE_NAME)

	// загружаем .env file
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("ошибка чтения файла конфигурации (.env file) key:", key)
	}

	return os.Getenv(key)
}
