package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wstiehler/tcc-backend/internal/environment"
)

func CORSMiddleware() gin.HandlerFunc {
	env := environment.GetInstance()
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", env.CORS_URL)
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, origin, Content-Type, accept")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Content-Type", "application/json")
		c.Header("User-Agent", "openvagas-service")

		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
