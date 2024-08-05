package infra

import (
	"github.com/gin-gonic/gin"
)

func AccountsRouter(context *gin.Context) {
	response_status := 200
	response_data := gin.H{"message": "tasks"}
	context.JSON(response_status, response_data)
}	