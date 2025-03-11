package handlers

import (
	"backend-image-service/models"
	"backend-image-service/services"
	"encoding/json"
	"net/http"
	"strconv"
)

var jobService = services.NewJobService()

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Count != len(req.Visits) {
		http.Error(w, "count does not match number of visits", http.StatusBadRequest)
		return
	}

	jobID := jobService.CreateJob(req)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"job_id": jobID})
}

func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	jobIDStr := r.URL.Query().Get("jobid")
	if jobIDStr == "" {
		http.Error(w, "Missing jobid", http.StatusBadRequest)
		return
	}

	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		http.Error(w, "Invalid jobid", http.StatusBadRequest)
		return
	}

	job, exists := jobService.GetJobStatus(jobID)
	if !exists {
		http.Error(w, "Job not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(job)
}
