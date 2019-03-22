package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jpsiyu/webapp-go/server/conf"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	indexPath := "dist/index.html"
	data, err := ioutil.ReadFile(indexPath)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	} else {
		w.Write(data)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.WriteHeader(404)
	w.Write([]byte("404 - " + http.StatusText(404)))
}

func main() {
	r := mux.NewRouter()
	// set handler
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("dist/"))))
	r.HandleFunc("/", homeHandler)
	r.NotFoundHandler = http.HandlerFunc(homeHandler)

	log.Println(fmt.Sprintf("Server listening on port %d", conf.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)
}
