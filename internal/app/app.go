package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/danmory/web-hackers-service/internal/storage/redis"
	appHTTP "github.com/danmory/web-hackers-service/internal/transport/http"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := createServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		address := fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
		if err := app.Listen(address); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-c
	if err := app.Shutdown(); err != nil {
		log.Panic(err)
	}
	if err := redis.CloseDB(); err != nil {
		log.Panic(err)
	}
}

func createServer() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Get("/json/hackers", appHTTP.GetHackers)
	return app
}
