package main

import (
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-commonmark/markdown"
)

var md *markdown.Markdown

func main() {
	// Create router (release)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Delims("{{", "}}")

	// TODO: add options
	md = markdown.New(markdown.HTML(true))

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
			"title":  "The Fifth Site - Answering your questions about 5G.",
			"points": points,
		})
	})
	// Set /assets/:file route
	r.GET("/assets/:file", func(c *gin.Context) {
		c.File(path.Join("assets", c.Param("file")))
	})

	// Start server
	chk(r.Run(":80"))
}

func chk(e error) {
	if e != nil {
		panic(e)
	}
}
