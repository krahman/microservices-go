package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type responseData struct {
	Message string `json:"message"`
}

type requestData struct {
	Name string `json:"name"`
}

const port = 8000

func main() {
	catHandler()
	server()
}

func catHandler() {
	// To serve a directory on disk (/tmp) under an alternate URL
	// path (/tmpfiles/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	tmpFileServer := http.FileServer(http.Dir("./tmp"))
	tmpFiles := "/tmpfiles/"
	http.Handle(tmpFiles, http.StripPrefix(tmpFiles, tmpFileServer))
}

func server() {
	http.HandleFunc("/", responseHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func responseHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var request requestData
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := responseData{ Message: "Hello World, " +  request.Name}


	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}