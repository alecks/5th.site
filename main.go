package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golang-commonmark/markdown"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

var md *markdown.Markdown

func main() {
	// Create router (debug)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Delims("{{", "}}")

	// TODO: add options
	md = markdown.New()

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
	r.GET("/", func(c *gin.Context) {
		points := getPoints()

		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"title":  "5G",
			"points": points,
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
