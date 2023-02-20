package util

import "net/http"

type ClientResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
	Message    string `json:"message"`
}

func SendApiSuccess(data any, message string) *ClientResponse {

	return buildSendData(true, http.StatusOK, data, message)
}

func SendApiError(code int, message string) *ClientResponse {

	return buildSendData(false, code, nil, message)
}

func buildSendData(success bool, code int, data any, message string) *ClientResponse {

	return &ClientResponse{
		Success:    success,
		StatusCode: code,
		Data:       data,
		Message:    message,
	}
}
