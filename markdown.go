package main

import (
	"html/template"
	"io/ioutil"
	"path"
	"sort"
	"strconv"
	"strings"
)

type byFname []string

func (s byFname) Len() int {
	return len(s)
}

func (s byFname) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byFname) Less(i, j int) bool {
	inds, err := strconv.Atoi(strings.Split(s[i], "_")[0])
	indj, err := strconv.Atoi(strings.Split(s[j], "_")[0])
	chk(err)
	return inds < indj
}

func getPoints() map[string]template.HTML {
	dir, err := ioutil.ReadDir("points")
	chk(err)

	preSort := map[string]template.HTML{}

	for _, v := range dir {
		name := v.Name()
		if name != "example.md" && strings.HasSuffix(name, ".md") {
			bytes, err := ioutil.ReadFile(path.Join("points", name))
			chk(err)

			preSort[name] = template.HTML(md.RenderToString(bytes))
		}
	}

	var keys []string
	for k := range preSort {
		keys = append(keys, k)
	}
	sort.Sort(byFname(keys))

	res := map[string]template.HTML{}
	for _, k := range keys {
		val := preSort[k]
		res[k] = val
	}

	return res
}
