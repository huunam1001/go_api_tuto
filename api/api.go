package api

import (
	db "go_api_tuto/db/sqlc"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	mongo  *mongo.Client
}

func NewServer(store db.Store, mongoDb *mongo.Client) Server {

	sever := &Server{
		store: store,
		mongo: mongoDb,
	}

	router := gin.Default()

	nonAuthGroup := router.Group("api/v1")
	{
		nonAuthGroup.POST("/user/register", sever.UserRegister)
		nonAuthGroup.POST("/user/login", sever.UserLogin)
	}

	authGroup := router.Group("api/v1")
	{
		authGroup.POST("/account", sever.CreateAccount)
		authGroup.GET("/user/me", sever.GetMe)
	}

	// router.GET("/user/me", sever.GetMe)

	sever.router = router

	return *sever
}

func (server *Server) Start(addres string) error {

	return server.router.Run(addres)
}
