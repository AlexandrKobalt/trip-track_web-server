package handler

import (
	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	Upload() fiber.Handler
	GetURL() fiber.Handler
}
