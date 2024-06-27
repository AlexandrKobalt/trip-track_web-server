package service

import (
	"context"

	"github.com/AlexandrKobalt/trip-track_web-server/internal/file/models"
)

type IService interface {
	Upload(
		ctx context.Context,
		params models.UploadParams,
	) (result models.UploadResult, err error)
	GetURL(
		ctx context.Context,
		params models.GetURLParams,
	) (result models.GetURLResult, err error)
}
