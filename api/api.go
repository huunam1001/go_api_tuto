package api

import (
	db "go_api_tuto/db/sqlc"

	"go_api_tuto/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	mongo  *mongo.Client
	redis  *redis.Client
}

func NewServer(store db.Store, mongoDb *mongo.Client, redis *redis.Client) Server {

	sever := &Server{
		store: store,
		mongo: mongoDb,
		redis: redis,
	}

	router := gin.Default()

	nonAuthGroup := router.Group(util.API_GROUPING)
	{
		nonAuthGroup.POST("/user/register", sever.UserRegister)
		nonAuthGroup.POST("/user/login", sever.UserLogin)
	}

	authGroup := router.Group(util.API_GROUPING, util.HeaderCheck(sever.redis, sever.mongo))
	{
		/// add middle ware here

		authGroup.POST("/account", sever.CreateAccount)
		authGroup.GET("/user/me", sever.GetMe)

		authGroup.GET("/category/get_list", sever.GetListCategory)
		authGroup.POST("/category/create", sever.AddCategory)
		authGroup.POST("/category/update", sever.UpdateCategory)
		authGroup.POST("/category/delete", sever.DeleteCategory)

		authGroup.GET("/product/get_list", sever.GetListProduct)
		authGroup.POST("/product/create", sever.AddProduct)
	}

	sever.router = router

	return *sever
}

func (server *Server) Start(addres string) error {

	return server.router.Run(addres)
}
