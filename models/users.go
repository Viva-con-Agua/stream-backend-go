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
	User struct {
		Uuid      string `json:"uuid" validate:"required"`
		Email     string `json:"email" validate:"required"`
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Roles     []Role `json:"roles" validate:"required"`
		Updated   int    `json:"updated" validate:"required"`
		Created   int    `json:"created" validate:"required"`
	}
	UserRole struct {
		UserId string `json:"user_uuid" validate:"required"`
		RoleId string `json:"role_uuid" validate:"required"`
	}
)
