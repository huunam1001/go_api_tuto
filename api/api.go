package api

import (
	db "BankTuto/db/sqlc"

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

	router.Group("api")
	router.POST("/account", sever.CreateAccount)

	sever.router = router

	return *sever
}

func (server *Server) Start(addres string) error {

	return server.router.Run(addres)
}

func Log() {
	////
	////
}

func errorResponse(e error) gin.H {
	return gin.H{"Error": e.Error()}
}
