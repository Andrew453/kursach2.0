package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"/home/prokiprok/go/src/kursach/html/home_page.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		app.notFound(w)
		return
	}

	files := []string{
		"/home/prokiprok/go/src/kursach/html/about.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) sources(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sources" {
		app.notFound(w)
		return
	}

	files := []string{
		"/home/prokiprok/go/src/kursach/html/sources.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
	}
}
