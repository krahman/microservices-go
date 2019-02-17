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

func main() {
	port := 8000

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