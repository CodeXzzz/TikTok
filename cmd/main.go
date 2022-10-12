package main

import (
	"Gin_GoBlog/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin
	r := gin.Default()
	router.InitRouter(r)

}
