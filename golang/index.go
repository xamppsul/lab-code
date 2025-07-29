package main

import (
	"golang/config"
	"log"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	// var num int
	// //array dulu
	// basic.Array()
	// //condition
	// fmt.Print("Masukkan angka: ")
	// fmt.Scan(&num) //input dari user

	// result := basic.ValidateNumberAvailable(num)
	// message := fmt.Sprintf("Angka %d, %s", basic.ExecNumber(num), result)
	// fmt.Println(message)

	//config env load and hot reload
	config.EnvConfig()
	fastergoding.Run()
	//hot reload jika env type tidak available
	if os.Getenv("APP_ENVIRONTMENT_TYPE") != "stagging" && os.Getenv("APP_ENVIRONTMENT_TYPE") != "production" {
		log.Fatal("Env Type wrong")
	}

	app := fiber.New()
	app.Get("/", func(res *fiber.Ctx) error {
		return res.SendString(runtime.Version() + " With fiber framework: " + os.Getenv("APP_DESCRIPTION"))
	})

	app.Listen(":" + os.Getenv("APP_PORT"))
}
