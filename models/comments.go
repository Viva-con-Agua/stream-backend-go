package models

type (
	CommentCreate struct {
		User User   `json:"user"`
		Text string `json:"text"`
		Tag  string `json:"tag" validate:"required"`
	}
	Comment struct {
		Uuid    string `json:"uuid" validate:"required"`
		User    User   `json:"user" validate:"required"`
		Text    string `json:"text"`
		Tag     string `json:"tag" validate:"required"`
		Created int64  `json:"created" validate:"required"`
	}
)
