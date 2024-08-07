package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ServerSentEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	for {
		// Write the event data
		fmt.Fprintf(w, "data: %s\n\n", time.Now().String())

		// Flush the data to the client
		flusher.Flush()

		// Sleep for a while before sending the next event
		time.Sleep(1 * time.Second)
	}

}
func main() {
	fmt.Println("Welcome to server sent events.")
	r := mux.NewRouter()
	r.HandleFunc("/sse", ServerSentEventsHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
