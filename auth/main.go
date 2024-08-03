package main

import (
	"github.com/gin-gonic/gin"
	routes "auth/infra/routes"
)

func main() {
	r := gin.Default()
	r.GET("/ping", routes.TasksRouter)
	r.Run()
}