package main

import (
	"html/template"
	"io/ioutil"
	"path"
	"strings"
)

func getPoints() []template.HTML {
	dir, err := ioutil.ReadDir("points")
	chk(err)

	var res []template.HTML

	for _, v := range dir {
		name := v.Name()
		if strings.HasSuffix(name, ".md") {
			bytes, err := ioutil.ReadFile(path.Join("points", name))
			chk(err)

			res = append(res, template.HTML(md.RenderToString(bytes)))
		}
	}

	return res
}
