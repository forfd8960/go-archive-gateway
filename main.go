package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("asset/*.html")
}
