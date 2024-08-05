package main

import (
	"github.com/gin-gonic/gin"
	routes "auth/infra/routes/accounts"
)

func main() {
	r := gin.Default()
	r.POST("/accounts", routes.CreateAccount)
	r.Run()
}