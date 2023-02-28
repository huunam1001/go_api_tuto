package api

import (
	db "go_api_tuto/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) Server {

	sever := &Server{
		store: store,
	}

	router := gin.Default()

	router.POST("/account", sever.CreateAccount)
	router.POST("/user/login", sever.UserLogin)

	sever.router = router

	return *sever
}

func (server *Server) Start(addres string) error {

	return server.router.Run(addres)
}
