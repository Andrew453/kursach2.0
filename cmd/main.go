package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":7777", "Web server address")

	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFa\t", log.Ldate|log.Ltime)
	errorLog := log.New(f, "ERRORы\t", log.Ldate|log.Ldate)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	fileServer := http.FileServer(http.Dir("./css"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/sources", app.about)

	mux.Handle("/css/", http.StripPrefix("/css", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Println("Start server 127.0.0.1", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
