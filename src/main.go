package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
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
	handler := newValidationHandler(newMessageHandler())

	http.Handle("/", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request requestData
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	h.next.ServeHTTP(rw, r)
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

type messageHandler struct {}

func newMessageHandler() http.Handler {
	return messageHandler{}
}

func (h messageHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := responseData{Message: "hello"}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1* time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error", err)
	}
}