package main

import (
	"auth/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", router.Router)
	r.Run()
}