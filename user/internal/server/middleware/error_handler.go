package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, err interface{}) {
	message := fmt.Sprintf("Internal server error: %v", err)
	resp := map[string]interface{}{
		"status":  1,
		"message": message,
	}
	ctx.JSON(http.StatusInternalServerError, resp)
}
