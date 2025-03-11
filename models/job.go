package models

type Job struct {
	JobID  int    `json:"job_id"`
	Status string `json:"status"`
	Error  []struct {
		StoreID string `json:"store_id"`
		Error   string `json:"error"`
	} `json:"error,omitempty"`
}

type SubmitRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}
