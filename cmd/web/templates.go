package main

import (
	"html/template"
	"net/http"
	"time"
)

type pageData map[string]any

var functions = template.FuncMap{
	"humanDate": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.UTC().Format("02 Jan 2006 at 15:04")
	},
}

func render(w http.ResponseWriter, r *http.Request, page string, data pageData) {
	t, err := template.New(page).Funcs(functions).ParseFiles("./assets/templates/"+page, "./assets/templates/main.layout.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, data)
}
