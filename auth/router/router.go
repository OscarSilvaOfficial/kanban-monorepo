package router

import (
	"github.com/gin-gonic/gin"
)

func Router(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
}