package api

import (
	"go_api_tuto/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) UserLogin(ctx *gin.Context) {

	util.SendApiError(ctx, http.StatusBadRequest, "FUCK U")
}
