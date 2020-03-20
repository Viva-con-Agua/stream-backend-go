package models

import "strconv"

type (
	CrewCreate struct {
		Name   LoginInfo `json:"loginInfo" validate:"required"`
		Cities []City    `json:"cities" validate:"required"`
	}
	Crew struct {
		Uuid    string `json:"uuid" validate:"required"`
		Name    string `json:"name" validate:"required"`
		Cities  []City `json:"cities"`
		Updated int    `json:"updated" validate:"required"`
		Created int    `json:"created" validate:"required"`
	}
	QueryCrew struct {
		Offset string `query:"offset" default:"0"`
		Count  string `query:"count" default:"40"`
		Name   string `query:"name" default:"%"`
		Sort   string `query:"sort"`
		SortBy string `query:"sortby"`
	}
	FilterCrew struct {
		Name string
	}
)

func (q *QueryCrew) Defaults() {
	if q.Offset == "" {
		q.Offset = "0"
	}
	if q.Count == "" {
		q.Count = "40"
	}
	if q.Name == "" {
		q.Name = "%"
	}
}

func (q *QueryCrew) Page() *Page {
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

func (q *QueryCrew) OrderBy() string {
	var asc = "ASC"
	if q.Sort == "DESC" {
		asc = " DESC"
	}
	var sort = "ORDER BY "
	if q.SortBy == "" {
		return ""
	}
	if q.SortBy == "name" {
		return sort + " p.name " + asc
	}
	return sort
}

func (q *QueryCrew) Filter() *FilterCrew {
	filter := new(FilterCrew)
	if q.Name != "" {
		filter.Name = q.Name
	} else {
		filter.Name = "%"
	}
	return filter
}
