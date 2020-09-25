package models

type HttpError struct {
	Message    string `json:"message"`
	HttpStatus int    `json:"status"`
}

