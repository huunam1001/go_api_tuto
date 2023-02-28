package util

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func HeaderCheck() gin.HandlerFunc {

	return func(c *gin.Context) {

		authToken := c.Request.Header.Get("Authorization")

		if authToken == "" {
			SendApiError(c, ERROR_LOGIN_TOKEN, "Error login token")
			c.Abort()

			return
		}

		splitToken := strings.Split(authToken, "Bearer ")

		if len(splitToken) != 2 {

			SendApiError(c, ERROR_LOGIN_TOKEN, "Error login token")
			c.Abort()

			return
		}

		// print(splitToken[1])

		c.Next()
	}
}
