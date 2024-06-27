package handler

import (
	"github.com/AlexandrKobalt/trip-track_web-server/internal/file/models"
	"github.com/AlexandrKobalt/trip-track_web-server/internal/file/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.IService
}

func New(service service.IService) IHandler {
	return &handler{service: service}
}

func (h *handler) Upload() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return err
		}

		params := models.UploadParams{File: fileHeader}

		result, err := h.service.Upload(c.Context(), params)
		if err != nil {
			return err
		}

		return c.JSON(result)
	}
}

func (h *handler) GetURL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := models.GetURLParams{
			Key: c.Query("key"),
		}

		result, err := h.service.GetURL(c.Context(), params)
		if err != nil {
			return err
		}

		return c.JSON(result)
	}
}
