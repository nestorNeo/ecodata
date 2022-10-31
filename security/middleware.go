package security

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer"

func (val *Validator) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "missing authorization info",
			})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		fmt.Println(tokenString)

		claims, err := val.IsValid(strings.TrimSpace(tokenString))

		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		} else {
			c.Set("CLAIMS", claims)
			c.Next()
		}
	}
}
