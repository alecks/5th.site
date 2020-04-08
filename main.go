package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

func main() {
	// Create router (debug)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Delims("{{", "}}")

	dir, err := ioutil.ReadDir("views")
	chk(err)

	var files []string
	for _, v := range dir {
		name := v.Name()
		if strings.HasSuffix(name, ".html") {
			files = append(files, path.Join("views", name))
		}
	}
	r.LoadHTMLFiles(files...)

	// Set / route
	// TODO: add an easy way to add points; possibly markdown
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"title": "5G",
			"points": map[string]string{
				"test": "testy",
			},
		})
	})

	// Start server
	chk(r.Run(":80"))
}

func chk(e error) {
	if e != nil {
		panic(e)
	}
}
