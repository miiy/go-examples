package main

import (
	"html/template"
	"regexp"
)

const lenPath := len("/view/")

var titleValidator = regexp.MustCompile("^/[a-zA-Z0-9]+$")
var templates = make(map[string]*template.Template)
var err error

type Page struct {
	Title string
	Body []byte
}

func int()  {
	for _, tmpl := range []string{"edit", "view"} {
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html"))
	}
}
