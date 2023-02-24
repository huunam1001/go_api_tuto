package api

import (
	db "go_api_tuto/db/sqlc"
	"go_api_tuto/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {

	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		util.SendApiError(ctx, http.StatusBadRequest, "Data validation")

		return
	}

	println("XXXXXX")
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	println("YYYYY")

	account, err := server.store.CreateAccount(ctx, arg)

	println("ZZZZ")

	if err != nil {
		util.SendApiError(ctx, http.StatusInternalServerError, "Internal server error")

		return
	}

	util.SendApiSuccess(ctx, account, "")
}
