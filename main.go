package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanay85/go-ecommerce-product/configs"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	configs.GetAbsPath()
	configs.ConnectToDB()

	r.Run(":3000")
}
