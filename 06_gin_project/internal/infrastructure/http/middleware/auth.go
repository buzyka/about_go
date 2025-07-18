package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AuthHeader = "Authorization"
	AuthToken = "KFC-1234567890"
)	

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Here you can implement your authentication logic
		// For example, check for a token in the header
		token := ctx.GetHeader(AuthHeader)
		if token != "Bearer "+AuthToken {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// If the token is valid, proceed to the next handler
		ctx.Next()
	}
}
