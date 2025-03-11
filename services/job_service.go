package services

import (
	"backend-image-service/models"
	"backend-image-service/utils"
	"sync"
)

type JobService struct {
	jobs     map[int]*models.Job
	jobMutex sync.Mutex
	nextJobID int
}

func NewJobService() *JobService {
	return &JobService{
		jobs:      make(map[int]*models.Job),
		nextJobID: 1,
	}
}

func (js *JobService) CreateJob(req models.SubmitRequest) int {
	js.jobMutex.Lock()
	jobID := js.nextJobID
	js.jobs[jobID] = &models.Job{JobID: jobID, Status: "ongoing"}
	js.nextJobID++
	js.jobMutex.Unlock()

	go js.processJob(jobID, req)

	return jobID
}

func (js *JobService) processJob(jobID int, req models.SubmitRequest) {
	for _, visit := range req.Visits {
		for _, imgURL := range visit.ImageURLs {
			if err := utils.ProcessImage(imgURL); err != nil {
				js.jobMutex.Lock()
				js.jobs[jobID].Status = "failed"
				js.jobs[jobID].Error = append(js.jobs[jobID].Error, struct {
					StoreID string `json:"store_id"`
					Error   string `json:"error"`
				}{visit.StoreID, err.Error()})
				js.jobMutex.Unlock()
			}
		}
	}

	js.jobMutex.Lock()
	if js.jobs[jobID].Status != "failed" {
		js.jobs[jobID].Status = "completed"
	}
	js.jobMutex.Unlock()
}

func (js *JobService) GetJobStatus(jobID int) (*models.Job, bool) {
	js.jobMutex.Lock()
	defer js.jobMutex.Unlock()
	job, exists := js.jobs[jobID]
	return job, exists
}
