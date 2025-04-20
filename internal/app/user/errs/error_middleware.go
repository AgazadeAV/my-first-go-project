package errs

import (
	"errors"
	"fmt"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors.Last().Err

		var fieldErrors validation.FieldErrors
		if errors.As(err, &fieldErrors) {
			respondWithError(ctx, http.StatusBadRequest, "Validation failed", fieldErrors)
			return
		}

		switch {
		case errors.Is(err, ErrBadUUID):
			respondWithError(ctx, http.StatusBadRequest, "Invalid UUID", nil)
		case errors.Is(err, ErrBadJSON):
			respondWithError(ctx, http.StatusBadRequest, "Invalid input", nil)
		case errors.Is(err, ErrNotFound):
			respondWithError(ctx, http.StatusNotFound, "User not found", nil)
		default:
			respondWithError(ctx, http.StatusInternalServerError, "Something went wrong", nil)
		}
	}
}

func respondWithError(ctx *gin.Context, status int, message string, details map[string]string) {
	ctx.JSON(status, ErrorResponse{
		Timestamp: time.Now(),
		Status:    fmt.Sprintf("%d %s", status, http.StatusText(status)),
		Message:   message,
		Errors:    details,
	})
}
