package main

import (
	"html/template"
	"io/ioutil"
	"path"
	"strings"
)

func getPoints() map[string]template.HTML {
	dir, err := ioutil.ReadDir("points")
	chk(err)

	res := map[string]template.HTML{}

	for _, v := range dir {
		name := v.Name()
		if name != "example.md" && strings.HasSuffix(name, ".md") {
			bytes, err := ioutil.ReadFile(path.Join("points", name))
			chk(err)

			res[name] = template.HTML(md.RenderToString(bytes))
		}
	}

	return res
}
