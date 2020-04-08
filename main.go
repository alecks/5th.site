package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create router (debug)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// Set / route
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "200/OK")
	})

	// Start server
	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}
