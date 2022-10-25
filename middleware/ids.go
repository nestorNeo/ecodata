package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Middleware injects or perform an action before actually processing the request
func GuidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Set("uuid", uuid.String())
		fmt.Printf("The request with uuid %s is started \n", uuid)
		c.Next()
		fmt.Printf("The request with uuid %s is served \n", uuid)
	}
}

func TempStorage(prefix, tmpdir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("SERVER_STAGING_TMP_FILES", tmpdir)
		c.Set("SERVER_PREFIX_TMP_FILES", prefix)
		c.Next()
	}
}
