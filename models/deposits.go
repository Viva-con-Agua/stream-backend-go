package models

import "strconv"

type (
	Deposit struct {
		Uuid          string        `json:"publicId" validate:"required"`
		Full          Amount        `json:"full" validate:"required"`
		Amount        DepositAmount `json:"amount" validate:"required"`
		Confirmed     Confirmed     `json:"confirmed" validate:"required"`
		Crew          Crew          `json:"crew" validate:"required"`
		Supporter     Supporter     `json:"supporter" validate:"required"`
		DateOfDeposit string        `json:"dateOfDeposit" validate:"required"`
		Updated       int           `json:"updated" validate:"required"`
		Created       int           `json:"created" validate:"required"`
	}

	DepositAmount struct {
		PublicId    string    `json:"publicId" validate:"required"`
		TakingId    string    `json:"takingId" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Confirmed   Confirmed `json:"confirmed" validate:"required"`
		Amount      Amount    `json:"amount" validate:"required"`
		Created     int       `json:"created" validate:"required"`
	}

	// TODO: Finish struct
	DepositConfirm struct {
		PublicId      string    `json:"publicId" validate:"required"`
		TakingId      string    `json:"takingId" validate:"required"`
		Description   string    `json:"description" validate:"required"`
		Confirmed     Confirmed `json:"confirmed" validate:"required"`
		DateOfDeposit string    `json:"dateOfDeposit" validate:"required"`
		Amount        Amount    `json:"amount" validate:"required"`
		Created       int       `json:"created" validate:"required"`
	}

	// TODO: Finish struct
	DepositCreate struct {
		PublicId      string    `json:"publicId" validate:"required"`
		TakingId      string    `json:"takingId" validate:"required"`
		Description   string    `json:"description" validate:"required"`
		DateOfDeposit string    `json:"dateOfDeposit" validate:"required"`
		Confirmed     Confirmed `json:"confirmed" validate:"required"`
		Amount        Amount    `json:"amount" validate:"required"`
		Created       int       `json:"created" validate:"required"`
	}

	QueryDeposit struct {
		Offset string `query:"offset" default:"0"`
		Count  string `query:"count" default:"50"`
		Sort   string `query:"sort"`
		SortBy string `query:"sortby"`
	}
	FilterDeposit struct {
		Name string
	}
)

func (q *QueryDeposit) Defaults() {
	if q.Offset == "" {
		q.Offset = "0"
	}
	if q.Count == "" {
		q.Count = "50"
	}
}

func (q *QueryDeposit) Page() *Page {
	var err error
	page := new(Page)
	page.Offset, err = strconv.Atoi(q.Offset)
	if err != nil {
		page.Offset = 0
	}
	page.Count, err = strconv.Atoi(q.Count)
	if err != nil {
		page.Count = 50
	}
	return page
}

func (q *QueryDeposit) OrderBy() string {
	var asc = "ASC"
	if q.Sort == "DESC" {
		asc = " DESC"
	}
	var sort = "ORDER BY "
	if q.SortBy == "" {
		return ""
	}
	return sort + " " + asc
}

func (q *QueryDeposit) Filter() *FilterDeposit {
	filter := new(FilterDeposit)
	return filter
}
