package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello-world", Handlers().HelloWorld)
	http.HandleFunc("/hello-world-html", Handlers().HelloWorldHTML)
	http.HandleFunc("/hello-world-json", Handlers().HelloWorldJSON)
	http.HandleFunc("/syllabus", SyllabusHandler().GetSyllabus)
	http.HandleFunc("/syllabus/delete", SyllabusHandler().DeleteSyllabus)
	http.HandleFunc("/syllabus/create", SyllabusHandler().CreateSyllabus)
	http.HandleFunc("/syllabus/update", SyllabusHandler().UpdateSyllabus)
	http.HandleFunc("/help", HelpHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	helpText := `
	Available APIs:
	- /hello-world: Responds with "Hello World - GWS"
	- /hello-world-html: Responds with HTML "Hello World - GWS" with H1 and mDash
	- /hello-world-json: Responds with JSON {"message": "Hello World - GWS"}
	- /syllabus: Responds with the embedded JSON syllabus
	- /syllabus/delete: Responds with "deleted - stubbed"
	- /syllabus/create: Responds with "create-stubbed"
	- /syllabus/update: Responds with "update-stubbed"
	`
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(helpText))
}
