package models

import "log"

type (
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
	log.Print(q.Email)
	if q.Email != "" {
		filter = filter + " Profile.email LIKE '" + q.Email + "'"
	}
	if filter == "WHERE" {
		return ""
	}
	return filter
}
