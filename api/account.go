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

		ctx.JSON(http.StatusOK, util.SendApiError(http.StatusBadRequest, "Data validation"))

		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusOK, util.SendApiError(http.StatusInternalServerError, "Internal server error"))

		return
	}

	ctx.JSON(http.StatusOK, util.SendApiSuccess(account, ""))
}
