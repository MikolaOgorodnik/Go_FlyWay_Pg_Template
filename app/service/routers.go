package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "It works!")
}
