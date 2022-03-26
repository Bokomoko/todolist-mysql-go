package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)
// *variables in go are pointers dereferrencing the value
// the following function receives a request and a respondeWriter and writes a response
// ResponseWriter is an interface that defines the methods that a responseWriter should implement
func Healthz(res http.ResponseWriter, req *http.Request) {

	log.Info("API Health is OK") // writes to log file
	res.Header().Set("Content-Type", "application/json") // set response header, defines response type
	io.WriteString(res, `{"alive": true}`) // write response body
}

// the init() function is called before the main function !!
// it can be used to instantiate variables and perform other initialization tasks
// it is called only once in the package lifetime
func init() {
	log.SetFormatter(&log.TextFormatter{}) // defines which function will carry on the formating of the logged message
	log.SetReportCaller(true)		// defines if the line number of the log message should be included
}

// the main function is the entry point of the program
// it´s where the program starts running
// it is called as many times as the program is run
func main() {
	log.Info("Starting Todolist API server")	// logs the start of the server
	router := mux.NewRouter()				// creates a new router
	router.HandleFunc("/healthz", Healthz).Methods("GET")	// defines the healthz endpoint
	http.ListenAndServe(":8000", router)				// starts the server that uses the router previously created
}