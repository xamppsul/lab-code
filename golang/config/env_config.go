package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvConfig(path string) {
	/*
		GetWd() handler untuk membaca root dir project,
		dengan tujuan agar dapat melacak keberadaan dari file .env
	*/
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	/*
		jalankan lokasi .env file
		nilai dari path berupa paramater yang diisi,
		pada saat pemanggilan fungsi EnvConfig
	*/
	env_path := filepath.Join(dir, path)
	err = godotenv.Load(env_path)

	/*
		validate path .env jika tidak ditemukan maka
		menampilkan pesan error berupa \
		"Err: wrong path .env atau lokasi path .env salah" dan keluar dari program
	*/
	if err != nil {
		log.Fatal("Err: wrong path .env")
		os.Exit(1)
	}
}
