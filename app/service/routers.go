package service

import (
	"Pg_FW_Template/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Index(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "It works!")
}

func GetTemplate(c *gin.Context) {
	s := NewTemplateService(dao.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
