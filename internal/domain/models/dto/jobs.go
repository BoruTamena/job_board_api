package dto

import "time"

type JobStatus bool

type JobType string

const (
	Open    JobStatus = true
	Clossed JobStatus = false
)

const (
	Parmanent = "PARMANENT"
	Temporary = "TEMPORARY"
	PartTime  = "PARTTIME"
)

type Jobs struct {
	Id          int       `json:"id"`
	PostByID    int       `json:"post_by_id"`
	Title       string    `json:"title"`
	Type        int       `json:"job_type_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	DeadLine    time.Time `json:"dead_line"`
	IsActive    JobStatus `json:"is_active"`
}

type JobTypes struct {
	Id   int     `json:"id"`
	Type JobType `json:"job_type"`
}
