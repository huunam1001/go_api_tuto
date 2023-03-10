package api

import (
	"database/sql"
	"fmt"
	db "go_api_tuto/db/sqlc"
	"go_api_tuto/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (server *Server) UserRegister(ctx *gin.Context) {

	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		util.SendValidationError(ctx)

		return
	}

	getArr := db.GetListUserWithAccountOrEmailParams{
		Username: req.Username,
		Email:    req.Email,
	}

	users, err := server.store.GetListUserWithAccountOrEmail(ctx, getArr)

	if err != nil {
		util.SendInternalServerError(ctx)

		return
	}

	if len(users) > 0 {

		util.SendApiError(ctx, util.ACCOUNT_REGITERED, "Username or email was registered already")

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

	userMap := make(map[string]string)
	userMap["username"] = user.Username
	userMap["fullName"] = user.FullName
	userMap["email"] = user.Email

	util.SendApiSuccess(ctx, userMap, "")
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) UserLogin(ctx *gin.Context) {

	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		util.SendValidationError(ctx)

		return
	}

	arg := db.GetLoginUserParams{
		Username:       req.Username,
		Email:          req.Username,
		HashedPassword: req.Password,
	}

	user, err := server.store.GetLoginUser(ctx, arg)

	if err != nil {
		fmt.Println(err.Error())

		if err != sql.ErrNoRows {
			util.SendInternalServerError(ctx)
		} else {
			util.SendApiError(ctx, util.ACCOUNT_NOT_FOUND, "Please check your account and password")
		}

		return
	}

	userMap := make(map[string]string)
	userMap["username"] = user.Username
	userMap["fullName"] = user.FullName
	userMap["email"] = user.Email

	token, err := util.GenerateJWT(userMap)

	if err != nil || len(token) == 0 {

		util.SendInternalServerError(ctx)
	}

	reponseMap := make(map[string]any)
	reponseMap["user"] = userMap
	reponseMap["token"] = token
	reponseMap["loginDate"] = primitive.NewDateTimeFromTime(time.Now())

	util.SendApiSuccess(ctx, reponseMap, "")

	util.SaveLoginToken(server.redis, server.mongo, token, reponseMap)
}

func (server *Server) GetMe(ctx *gin.Context) {

	util.SendValidationError(ctx)
}
