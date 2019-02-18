package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type ResponseTest struct {
	Message string
}

func BenchmarkResponseHandlerVariable(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := ResponseTest{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(response)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkResponseHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := ResponseTest{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
	}
}

func BenchmarkResponseHandlerReference(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := ResponseTest{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}

func BenchmarkResponseHandler(b *testing.B) {
	b.ResetTimer()

	var body = bytes.NewBuffer([]byte(`"Name": "World"`))
	for i := 0; i < b.N; i++ {
		r, _ := http.Post(
			"http://localhost:8000/",
			"application/json",
			body,
		)

		var response responseData
		decoder := json.NewDecoder(r.Body)

		_ = decoder.Decode(&response)
	}
}

func init() {
	go server()
}