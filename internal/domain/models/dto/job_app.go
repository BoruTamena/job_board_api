package dto

import "time"

type Application struct {
	Id      int       `json:"id"`
	UserId  int       `json:"user_id"`
	PostId  int       `json:"job_post_id"`
	AppDate time.Time `json:"apply_date"`
}
