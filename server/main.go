package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func main() {
	// set handler
	fs := http.FileServer(http.Dir("dist/"))
	http.Handle("/", fs)

	// parse json file
	var result map[string]interface{}
	parsePackageJson(&result)
	port := result["port"]

	log.Println(fmt.Sprintf("Server listening on port %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
