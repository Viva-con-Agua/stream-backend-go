package models

type (
	LoginInfo struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	UserCreate struct {
		LoginInfo LoginInfo `json:"loginInfo" validate:"required"`
		FirstName string    `json:"first_name" validate:"required"`
		LastName  string    `json:"Last_name" validate:"required"`
	}
	UserSignIn struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
