package models

import "strconv"

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
		Count  string `query:"count" default:"40"`
		Email  string `query:"email" default:"%"`
		Sort   string `query:"sort"`
		SortBy string `query:"sortby"`
	}
)

func (q *QueryUser) Defaults() {
	if q.Offset == "" {
		q.Offset = "0"
	}
	if q.Count == "" {
		q.Count = "40"
	}
	if q.Email == "" {
		q.Email = "%"
	}
}

func (q *QueryUser) Page() *Page {
	var err error
	page := new(Page)
	page.Offset, err = strconv.Atoi(q.Offset)
	if err != nil {
		page.Offset = 0
	}
	page.Count, err = strconv.Atoi(q.Count)
	if err != nil {
		page.Count = 40
	}
	return page
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
