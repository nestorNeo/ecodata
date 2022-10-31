package security

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer"

func (val *Validator) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		claims, err := val.IsValid(tokenString)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Set("CLAIMS", claims)
			c.Next()
		}
	}
}
