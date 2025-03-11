package test

import (
	"backend-image-service/models"
	"backend-image-service/services"
	"testing"
)

// Test job creation
func TestCreateJob(t *testing.T) {
	jobService := services.NewJobService()

	request := models.SubmitRequest{
		Count: 1,
		Visits: []models.Visit{
			{
				StoreID:   "S00339218",
				ImageURLs: []string{"https://www.gstatic.com/webp/gallery/2.jpg"},
				VisitTime: "2025-03-11",
			},
		},
	}

	jobID := jobService.CreateJob(request)

	if jobID != 1 {
		t.Errorf("Expected job ID 1, got %d", jobID)
	}
}

// Test job status retrieval
func TestGetJobStatus(t *testing.T) {
	jobService := services.NewJobService()

	request := models.SubmitRequest{
		Count: 1,
		Visits: []models.Visit{
			{
				StoreID:   "S00339218",
				ImageURLs: []string{"https://www.gstatic.com/webp/gallery/2.jpg"},
				VisitTime: "2025-03-11",
			},
		},
	}

	jobID := jobService.CreateJob(request)

	job, exists := jobService.GetJobStatus(jobID)
	if !exists {
		t.Errorf("Expected job ID %d to exist", jobID)
	}

	if job.Status != "ongoing" {
		t.Errorf("Expected job status 'ongoing', got '%s'", job.Status)
	}
}
