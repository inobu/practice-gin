package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/inobu/practice-gin/application"
)

func bookRouter(router *gin.RouterGroup) {
	router.GET("/", application.GetBook)
}
