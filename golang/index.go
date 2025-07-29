package main

import (
	"golang/config"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	//load .env
	config.EnvConfig("../golang/.env")
	//hot reload
	fastergoding.Run()
	if err := config.SetTimezone(os.Getenv("GO_TIMEZONE")); err != nil {
		log.Fatal(err)
	}
	//print timezone base location
	time := config.GetTimezone(time.Now())
	log.Print("Asia Jakarta :" + time.Format(os.Getenv("GO_TIMEZONE_FORMAT")))
	//instance fiber load server
	app := fiber.New()
	app.Get("/", func(res *fiber.Ctx) error {
		return res.SendString(runtime.Version() + " With fiber framework: " + os.Getenv("GO_DESCRIPTION"))
	})
	app.Listen(":" + os.Getenv("GO_PORT"))
	//validate timezone location seting

}
