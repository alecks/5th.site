package main

import (
	"github.com/gin-gonic/gin"
	"path"
)

func main() {
	// Create router (debug)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// Set / route
	// TODO: template page
	r.GET("/", func(c *gin.Context) {
		c.File(path.Join("views", "index.html"))
	})

	// Start server
	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}
