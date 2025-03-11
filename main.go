package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"backend-image-service/handlers"
)

func main() {
	http.HandleFunc("/api/submit/", handlers.SubmitJobHandler)
	http.HandleFunc("/api/status", handlers.GetJobStatusHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
