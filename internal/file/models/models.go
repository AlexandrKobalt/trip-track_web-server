package models

import "mime/multipart"

type UploadParams struct {
	File *multipart.FileHeader
}

type UploadResult struct {
	Key string `json:"key"`
}

type GetURLParams struct {
	Key string `json:"key"`
}

type GetURLResult struct {
	URL string `json:"url"`
}
