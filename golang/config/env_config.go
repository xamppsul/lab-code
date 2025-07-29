package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvConfig(path string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	env_path := filepath.Join(dir, path)
	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal("Err: env doesn't exist")
		os.Exit(1)
	}
}
