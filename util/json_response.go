package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data,omitempty"`
	Message    string `json:"message"`
}

func SendApiSuccess(ctx *gin.Context, data any, message string) {

	ctx.JSON(http.StatusOK, buildSendData(true, http.StatusOK, data, message))
}

func SendApiError(ctx *gin.Context, code int, message string) {

	ctx.JSON(http.StatusOK, buildSendData(false, code, nil, message))
}

func SendValidationError(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, buildSendData(false, http.StatusBadRequest, nil, "Data validation"))
}

func SendInternalServerError(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, buildSendData(false, http.StatusInternalServerError, nil, "Internal server error"))
}

func buildSendData(success bool, code int, data any, message string) *ClientResponse {

	return &ClientResponse{
		Success:    success,
		StatusCode: code,
		Data:       data,
		Message:    message,
	}
}
