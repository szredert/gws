package main

import (
	"embed"
	"net/http"
)

var syllabusFile embed.FS

type SyllabusHandlerStruct struct{}

func SyllabusHandler() SyllabusHandlerStruct {
	return SyllabusHandlerStruct{}
}

func (s SyllabusHandlerStruct) GetSyllabus(w http.ResponseWriter, r *http.Request) {
	syllabusData, err := syllabusFile.ReadFile("syllabus.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error reading syllabus file"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(syllabusData)
}

func (s SyllabusHandlerStruct) DeleteSyllabus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("deleted - stubbed"))
}

func (s SyllabusHandlerStruct) CreateSyllabus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("create-stubbed"))
}

func (s SyllabusHandlerStruct) UpdateSyllabus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("update-stubbed"))
}
