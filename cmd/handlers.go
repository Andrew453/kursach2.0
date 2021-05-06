package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/h" {
		app.notFound(w)
		return
	}
	comms, err := app.comments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	file := "/home/f0rchy/go/src/kursach2.0/html/home_page.html"

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		app.errorLog.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "comments", comms)

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
		"/home/f0rchy/go/src/kursach2.0/html/about.html",
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
		"/home/f0rchy/go/src/kursach2.0/html/sources.html",
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
