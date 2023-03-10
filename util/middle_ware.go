package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type TokenLogin struct {
	Authorized bool      `json:"authorized"`
	Exp        int64     `json:"exp"`
	Data       LoginData `json:"data"`
}

type LoginData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

func HeaderCheck() gin.HandlerFunc {

	return func(c *gin.Context) {

		authToken := c.Request.Header.Get("Authorization")

		parts := strings.Split(authToken, " ")

		if parts[0] != "Bearer" ||
			len(parts) < 2 ||
			strings.TrimSpace(parts[1]) == "" {
			SendApiError(c, ERROR_LOGIN_TOKEN, "Error login token")

			c.Abort()

			return
		}

		tokenMap, success := ExtractClaims(parts[1])

		if !success {
			SendApiError(c, ERROR_LOGIN_TOKEN, "Error login token")

			c.Abort()

			return
		}

		jsonData, _ := json.Marshal(tokenMap)

		loginData := TokenLogin{}

		if err := json.Unmarshal(jsonData, &loginData); err != nil {

			SendInternalServerError(c)
			c.Abort()

			return
		}

		c.Request.Header.Add("username", loginData.Data.Username)
		c.Request.Header.Add("email", loginData.Data.Email)
		c.Request.Header.Add("fullName", loginData.Data.FullName)

		c.Next()
	}
}

func GetUserFromRequest(c *gin.Context) (*LoginData, bool) {

	userName := c.Request.Header.Get("username")
	email := c.Request.Header.Get("email")
	fullName := c.Request.Header.Get("fullName")

	if len(userName) > 0 {

		return &LoginData{
			Username: userName,
			Email:    email,
			FullName: fullName,
		}, true
	} else {
		return nil, false
	}
}

func SaveRedisToken(redis *redis.Client, key string, value interface{}) {

	if redis != nil {
		write := redis.Set(key, value, time.Duration(time.Hour*24*365*10))

		if write != nil {
			print(write.Err().Error())
		}
	}
}

func GetPagingFromRequest(c *gin.Context) (int, int) {
	getPage := c.Query("page")
	getLimit := c.Query("limit")

	page := 0
	limit := 10

	fmt.Sscan(getPage, &page)

	if page < 0 {
		page = 0
	}

	fmt.Sscan(getLimit, &limit)

	if limit < 0 {
		limit = 10
	}

	return page, limit
}
