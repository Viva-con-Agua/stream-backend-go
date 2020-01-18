package models

type (
	UserCreate struct {
		LoginInfo LoginInfo `json:"loginInfo" validate:"required"`
		FirstName string    `json:"first_name" validate:"required"`
		LastName  string    `json:"Last_name" validate:"required"`
	}
	User struct {
		Uuid      string `json:"uuid" validate:"required"`
		Email     string `json:"email" validate:"required"`
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Roles     []Role `json:"roles"`
		Updated   int    `json:"updated" validate:"required"`
		Created   int    `json:"created" validate:"required"`
	}
	QueryUser struct {
		Offset string `query:"offset" default:"0"`
		Size   string `query:"size" default:"40"`
		Email  string `query:"email" default:"%"`
		Sort   string `query:"sort"`
		SortBy string `query:"sortby"`
	}
)

func (q *QueryUser) Defaults() {
	if q.Offset == "" {
		q.Offset = "0"
	}
	if q.Size == "" {
		q.Size = "40"
	}
	if q.Email == "" {
		q.Email = "%"
	}
}

func (q *QueryUser) Page() string {
	if q.Size != "" {
		if q.Offset != "" {
			return "LIMIT " + q.Offset + ", " + q.Size
		}
		return "LIMIT " + q.Size
	}
	return ""
}

func (q *QueryUser) OrderBy() string {
	var asc = "ASC"
	if q.Sort == "DESC" {
		asc = " DESC"
	}
	var sort = "ORDER BY "
	if q.SortBy == "" {
		return ""
	}
	if q.SortBy == "email" {
		return sort + " Profile.email " + asc
	}
	return sort
}

func (q *QueryUser) Filter() string {
	var filter = "WHERE"
	if q.Email != "" {
		filter = filter + " Profile.email LIKE '" + q.Email + "'"
	}
	if filter == "WHERE" {
		return ""
	}
	return filter
}
