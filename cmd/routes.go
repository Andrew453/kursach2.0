package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/h", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/sources", app.sources)
	mux.HandleFunc("/commentform", func(w http.ResponseWriter, r *http.Request) {

		username := r.FormValue("username")
		comment := r.FormValue("text_comment")
		_, err := app.comments.Insert(username, comment)
		if err != nil {
			app.errorLog.Println(err)
		}

		http.Redirect(w, r, "/h#comments", http.StatusFound)
	})

	fileServer := http.FileServer(http.Dir("./css"))

	mux.Handle("/css/", http.StripPrefix("/css", fileServer))
	return mux
}
