package http

import (
	"github.com/danmory/web-hackers-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetHackers(ctx *fiber.Ctx) error {
	return ctx.JSON(service.RetrieveHackers())
}
