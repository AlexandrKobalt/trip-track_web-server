package service

import (
	"context"
	"io"

	fileserverproto "github.com/AlexandrKobalt/trip-track/backend/proto/fileserver"
	"github.com/AlexandrKobalt/trip-track_web-server/internal/file/models"
)

type service struct {
	fileClient fileserverproto.FileClient
}

func New(fileClient fileserverproto.FileClient) IService {
	return &service{fileClient: fileClient}
}

func (s *service) Upload(
	ctx context.Context,
	params models.UploadParams,
) (result models.UploadResult, err error) {
	file, err := params.File.Open()
	if err != nil {
		return result, err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return result, err
	}

	response, err := s.fileClient.Upload(
		ctx,
		&fileserverproto.UploadRequest{
			File: fileBytes,
		},
	)

	return models.UploadResult{
		Key: response.GetKey(),
	}, nil
}

func (s *service) GetURL(
	ctx context.Context,
	params models.GetURLParams,
) (result models.GetURLResult, err error) {
	response, err := s.fileClient.GetURL(
		ctx,
		&fileserverproto.GetURLRequest{
			Key: params.Key,
		},
	)
	if err != nil {
		return result, err
	}

	return models.GetURLResult{
		URL: response.GetUrl(),
	}, nil
}
