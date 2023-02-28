package api

import (
	db "go_api_tuto/db/sqlc"
	"go_api_tuto/util"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (server *Server) UserLogin(ctx *gin.Context) {

	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		util.SendValidationError(ctx)

		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.Password,
		Email:          req.Email,
		FullName:       req.FullName,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		util.SendInternalServerError(ctx)

		return
	}

	util.SendApiSuccess(ctx, user, "")
}
