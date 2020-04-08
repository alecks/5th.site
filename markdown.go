package main

import (
	"io/ioutil"
	"path"
	"strings"
)

func getPoints() []string {
	dir, err := ioutil.ReadDir("points")
	chk(err)

	var res []string

	for _, v := range dir {
		name := v.Name()
		if strings.HasSuffix(name, ".md") {
			bytes, err := ioutil.ReadFile(path.Join("points", name))
			chk(err)

			res = append(res, md.RenderToString(bytes))
		}
	}

	return res
}
