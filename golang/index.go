package main

import (
	"golang/config"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	//load .env
	config.EnvConfig("../golang/.env")
	//hot reload
	fastergoding.Run()

	app := fiber.New()
	app.Get("/", func(res *fiber.Ctx) error {
		return res.SendString(runtime.Version() + " With fiber framework: " + os.Getenv("APP_DESCRIPTION"))
	})

	app.Listen(":" + os.Getenv("APP_PORT"))
}
