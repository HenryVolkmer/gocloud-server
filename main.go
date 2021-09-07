package main

import(
	//"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")

}