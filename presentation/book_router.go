package presentation

import (
	"github.com/gin-gonic/gin"
)

func bookRouter(router *gin.RouterGroup) {
	router.GET("/:id", bookHandler)
}
