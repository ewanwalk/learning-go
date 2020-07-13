package main

import (
	"encoding/json"
	"net/http"
)

type Job struct {
	Id   string
	Name string
}

var jobs = map[string]Job{}

func Handler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		// Serve the resource.
		ind := request.URL.Path
		writer.Header().Set("Content-Type", "application/json")
		job := jobs[ind[1:]]
		json.NewEncoder(writer).Encode(job)
	case http.MethodPost:
		// Create a new job.
		var job Job
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&job)
		if err != nil {
			panic(err)
		}
		jobs[job.Id] = job
	default:
		// Give an error message.
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", mux)
}
