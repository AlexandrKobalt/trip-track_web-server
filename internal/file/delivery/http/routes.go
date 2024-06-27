package http

import (
	"github.com/AlexandrKobalt/trip-track_web-server/internal/file/handler"
	"github.com/gofiber/fiber/v2"
)

func Map(group fiber.Group, h handler.IHandler) {
	group.Post("/", h.Upload())
	group.Get("/", h.GetURL())
}
