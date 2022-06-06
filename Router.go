package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/addCryptoMessage", AddCryptoMessage)
	return r
}
