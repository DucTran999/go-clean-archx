// Package controller handles incoming HTTP requests and sends appropriate responses.
// It acts as the delivery layer in Clean Architecture, connecting HTTP routes to usecases.
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse defines the standard structure for API responses.
type APIResponse struct {
	Message string `json:"message,omitempty"` // Optional success or general message
	Error   string `json:"error,omitempty"`   // Optional error message
	Data    any    `json:"data,omitempty"`    // Optional payload data
}

// JSONResponse sends a structured JSON response with the given status code.
func JSONResponse(ctx *gin.Context, status int, res APIResponse) {
	ctx.JSON(status, res)
}

// JSONInternalErrorResponse sends a 500 Internal Server Error response with a safe message.
func JSONInternalErrorResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, APIResponse{
		Message: msg,
		Error:   http.StatusText(http.StatusInternalServerError),
	})
}

// JSONBadRequestResponse sends a 400 Bad Request response with the given error details.
func JSONBadRequestResponse(ctx *gin.Context, msg string, err error) {
	ctx.JSON(http.StatusBadRequest, APIResponse{
		Message: msg,
		Error:   err.Error(),
	})
}
