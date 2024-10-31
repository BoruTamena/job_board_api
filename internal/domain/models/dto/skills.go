package dto

type Skills struct {
	Id          int    `json:"id"`
	PostId      string `json:"job_post_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Level       string `json:"skill_level"`
}
