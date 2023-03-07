package api

import (
	db "go_api_tuto/db/sqlc"

	"go_api_tuto/util"

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

	nonAuthGroup := router.Group(util.API_GROUPING)
	{
		nonAuthGroup.POST("/user/register", sever.UserRegister)
		nonAuthGroup.POST("/user/login", sever.UserLogin)
	}

	authGroup := router.Group(util.API_GROUPING, util.HeaderCheck())
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

	// router.GET("/user/me", sever.GetMe)

	sever.router = router

	return *sever
}

func (server *Server) Start(addres string) error {

	return server.router.Run(addres)
}
