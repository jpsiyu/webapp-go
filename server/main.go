package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// parse package.json, read the application port
func parsePackageJson(result *map[string]interface{}) {
	pJson, err := os.Open("package.json")
	if err != nil {
		log.Println("Read package json error:", err)
		return
	}

	defer pJson.Close()
	byteValue, _ := ioutil.ReadAll(pJson)
	json.Unmarshal([]byte(byteValue), &result)
}

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

	// parse json file
	var result map[string]interface{}
	parsePackageJson(&result)
	port := result["port"]

	log.Println(fmt.Sprintf("Server listening on port %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
