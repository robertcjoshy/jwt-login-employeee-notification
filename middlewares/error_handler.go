package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e *AppError) Error() string { // converts the error to custom error of type apperror
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewAppError(code int, message string) *AppError { // converts the error to custom error of type apperror
	return &AppError{Code: code, Message: message}
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	// error handling -----------------------
	err := c.Errors.Last()

	if err != nil {
		var appErr *AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "an errror occured"})
		}
	}
}
