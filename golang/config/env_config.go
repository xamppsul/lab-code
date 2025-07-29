package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvConfig() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	env_path := filepath.Join(dir, "../golang/.env")
	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal("Err: env doesn't exist")
	}
}
