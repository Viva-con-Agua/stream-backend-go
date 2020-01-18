package models

type (
	LoginInfo struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	DeleteBody struct {
		Uuid string `json:"uuid" validate:"required"`
	}
	AssignBody struct {
		Assign string `json:"assign" validate:"required"`
		To     string `json:"to" validate:"required"`
	}
)
