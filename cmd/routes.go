package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/sources", app.sources)

	fileServer := http.FileServer(http.Dir("./css"))

	mux.Handle("/css/", http.StripPrefix("/css", fileServer))
	return mux
}
