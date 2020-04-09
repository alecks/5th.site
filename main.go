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

	md = markdown.New(markdown.HTML(true))

	dirs := []string{"views", "partials"}

	var files []string
	for _, d := range dirs {
		read, err := ioutil.ReadDir(d)
		chk(err)

		for _, v := range read {
			name := v.Name()
			if strings.HasSuffix(name, ".html") {
				files = append(files, path.Join(d, name))
			}
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
	// Handle 404
	// TODO: Handle all errors like this
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", map[string]interface{}{
			"errorCode":    "404",             // Use http.StatusNotFound in production
			"errorMessage": "Not Found",       // Removed as it would return same val: http.StatusText(http.StatusNotFound),
			"title":        "404 - Not Found", // Should be both http.StatusNotFound and http.StatusText concatinated
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
