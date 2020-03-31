package api

import (
	"net/http"

	"github.com/SergeyDavidenko/subscription/config"
	"github.com/SergeyDavidenko/subscription/db"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		token, err := c.Cookie(config.Conf.API.CokiesName)
		if err != nil {
			log.Debug(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Not found cookies, you not login",
			})
			return
		}
		if config.Conf.Redis.UseRedis {
			errValid := db.Client.CheckKey(token)
			if errValid != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": errValid.Error(),
				})
				return
			}
		}
		c.Next()
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}
