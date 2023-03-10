package util

import (
	"context"
	"encoding/json"
	"fmt"
	"go_api_tuto/db/mongodb"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func HeaderCheck(redis *redis.Client, mongoDb *mongo.Client) gin.HandlerFunc {

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

		var pass = false

		if redis != nil {
			val, err := redis.Get(parts[1]).Result()

			println(val)

			if err != nil {
				apiMongoDb := mongoDb.Database(MONGO_DATA_BASE)
				tokenCollection := apiMongoDb.Collection(MONGO_TOKEN_COLLECTION)

				filter := bson.D{{Key: "token", Value: parts[1]}}

				var result mongodb.LoginTokenInfo
				err = tokenCollection.FindOne(context.TODO(), filter).Decode(&result)

				if err == nil {
					/// Save token again to redis
					redis.Set(parts[1], parts[1], time.Duration(time.Hour*24*365*10))

					pass = true
				}

			} else {
				pass = true
			}
		}

		if !pass {

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

func SaveLoginToken(redis *redis.Client, mongoDb *mongo.Client, key string, value interface{}) {

	if redis != nil {
		redis.Set(key, key, time.Duration(time.Hour*24*365*10))
	}

	if mongoDb != nil {
		/// insert data to mongo atlas
		apiMongoDb := mongoDb.Database(MONGO_DATA_BASE)
		tokenCollection := apiMongoDb.Collection(MONGO_TOKEN_COLLECTION)

		tokenCollection.InsertOne(context.Background(), value)
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
