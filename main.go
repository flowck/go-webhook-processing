package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Job struct {
	Id         string `json:"id"`
	Status     string `json:"status"`
	WebhookUrl string `json:"webhookUrl"`
}

var jobList []Job

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func getJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := json.Marshal(jobList)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(jobs))
}

func createJob(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newJob Job
	err = json.Unmarshal(body, &newJob)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newJob.Id = uuid.NewString()
	newJob.Status = "unprocessed"
	jobList = append(jobList, newJob)
	jobResponse, err := json.Marshal(newJob)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jobResponse))
}

func jobHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getJobs(w, r)
		return
	case http.MethodPost:
		createJob(w, r)
		return
	}
}

func main() {
	Port := 4000
	Host := "localhost"
	SrvAddress := fmt.Sprintf("%s:%s", Host, strconv.Itoa(Port))

	fmt.Println("Server is up and running at", SrvAddress)

	// Handlers
	http.Handle("/jobs", middleware(http.HandlerFunc(jobHandlers)))

	err := http.ListenAndServe(SrvAddress, nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}
