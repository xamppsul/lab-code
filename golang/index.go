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
	config.EnvConfig(".env")
	//hot reload
	fastergoding.Run()
	//validate timezone location seting
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
	//load port for run server
	app.Listen(":" + os.Getenv("GO_PORT"))

}
