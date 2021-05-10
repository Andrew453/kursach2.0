package main

import (
	"html/template"
	"net/http"
)

// Обработка главной страницы
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/h" {
		app.notFound(w)
		return
	}
	comms, err := app.comments.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	file := "/home/prokiprok/go/src/kursach2.0/html/home_page.html"

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		app.errorLog.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "comments", comms)

	if err != nil {
		app.serverError(w, err)
	}
}

// Обработка страницы "О нас"
func (app *application) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		app.notFound(w)
		return
	}

	files := []string{
		"/home/prokiprok/go/src/kursach2.0/html/about.html",
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

// Обработка страницы "Источники"
func (app *application) sources(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sources" {
		app.notFound(w)
		return
	}

	files := []string{
		"/home/prokiprok/go/src/kursach2.0/html/sources.html",
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
