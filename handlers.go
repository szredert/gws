package main

import (
	"encoding/json"
	"net/http"
)

type HandlersStruct struct{}

func Handlers() HandlersStruct {
	return HandlersStruct{}
}

func (h HandlersStruct) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World - GWS"))
}

func (h HandlersStruct) HelloWorldHTML(w http.ResponseWriter, r *http.Request) {
	html := "<h1>Hello World - GWS</h1>"
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(html))
}

func (h HandlersStruct) HelloWorldJSON(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{"message": "Hello World - GWS"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
