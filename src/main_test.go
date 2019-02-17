package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	response := ResponseTest{ Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}