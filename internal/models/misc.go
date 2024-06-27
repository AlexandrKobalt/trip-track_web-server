package models

type BadRequestResponse struct {
	Cause string `json:"cause"`
}

type IDResponse struct {
	ID string `json:"id"`
}
