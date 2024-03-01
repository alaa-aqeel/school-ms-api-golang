package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alaa-aqeel/school-ms-api-golang/internal/util"
	"github.com/joho/godotenv"
)

func LoadConfig(envName string) {
	baseDir := util.GetBaseDir(".")
	err := godotenv.Load(filepath.Join(baseDir, envName))
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
}

func Env(key string, _default any) any {
	var env = os.Getenv(key)
	if env != "" {
		return env
	}
	return _default
}
