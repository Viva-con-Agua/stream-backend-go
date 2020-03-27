package models

import "strconv"

type (
    Taking struct {
        Uuid         string        `json:"id" validate:"required"`
        Amount       TakingAmount  `json:"amount" validate:"required"`
        Context      Context       `json:"context" validate:"required"`
        Comment      string        `json:"comment" validate:"required"`
        Details      Details       `json:"details" validate:"required"`
        DepositUnits []DepositUnit `json:"depositUnits" validate:"required"`
        Author       string        `json:"author" validate:"required"`
        Crew         Crew          `json:"crew" validate:"required"`
        Updated      int           `json:"updated" validate:"required"`
        Created      int           `json:"created" validate:"required"`
    }

    TakingAmount struct {
        Received  int         `json:"received" validate:"required"`
        Supporter []Supporter `json:"involvedSupporter" validate:"required"`
        Sources   []Source    `json:"sources" validate:"required"`
    }

    Supporter struct {
        Uuid string `json:"uuid" validate:"required"`
        Name string `json:"name" validate:"required"`
    }

    Source struct {
        Uuid         string       `json:"publicId" validate:"required"`
        Category     string       `json:"category" validate:"required"`
        Amount       Amount       `json:"amount" validate:"required"`
        TypeOfSource TypeOfSource `json:"typeOfSource" validate:"required"`
        Norms        string       `json:"norms" validate:"required"`
    }

    TypeOfSource struct {
        Category string   `json:"category" validate:"required"`
        External External `json:"external" validate:"required"`
    }

    Confirmed struct {
        Date int       `json:"date" validate:"required"`
        User Supporter `json:"user" validate:"required"`
    }

    External struct {
        location      string `json:"location" validate:"required"`
        contactPerson string `json:"contactPerson" validate:"required"`
        email         string `json:"email" validate:"required"`
        address       string `json:"address" validate:"required"`
        receipt       bool   `json:"receipt" validate:"required"`
    }

    Context struct {
        Description string `json:"description" validate:"required"`
        Category    string `json:"category" validate:"required"`
    }

    Details struct {
        ReasonForPayment string `json:"reasonForPayment" validate:"required"`
        Receipt          string `json:"receipt" validate:"required"`
    }

    DepositUnit struct {
        PublicId  string    `json:"publicId" validate:"required"`
        TakingId  string    `json:"takingId" validate:"required"`
        Confirmed Confirmed `json:"confirmed" validate:"required"`
        Amount    Amount    `json:"amount" validate:"required"`
        Created   int       `json:"created" validate:"required"`
    }

    Amount struct {
        Amount   int    `json:"amount" validate:"required"`
        Currency string `json:"currency" validate:"required"`
    }

    Crew struct {
        Uuid string `json:"uuid" validate:"required"`
        Name string `json:"name" validate:"required"`
    }

    // TODO: Finish struct
    TakingCreate struct {
        Email     string `json:"email" validate:"required"`
        FirstName string `json:"first_name" validate:"required"`
        LastName  string `json:"last_name" validate:"required"`
        Mobile    string `json:"birthdate" validate:"required"`
        Birthdate int    `json:"birthdate" validate:"required"`
        Sex       string `json:"sex" validate:"required"`
    }

    QueryTaking struct {
        Offset string `query:"offset" default:"0"`
        Count  string `query:"count" default:"50"`
        Name   string `query:"name" default:"%"`
        Crew   string `query:"crew" default:"%"`
        Sort   string `query:"sort"`
        SortBy string `query:"sortby"`
    }
    FilterTaking struct {
        Name string
        Crew string
    }
)

func (q *QueryTaking) Defaults() {
    if q.Offset == "" {
        q.Offset = "0"
    }
    if q.Count == "" {
        q.Count = "50"
    }
    if q.Name == "" {
        q.Name = "%"
    }
    if q.Crew == "" {
        q.Crew = "%"
    }
}

func (q *QueryTaking) Page() *Page {
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

func (q *QueryTaking) OrderBy() string {
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

func (q *QueryTaking) Filter() *FilterTaking {
    filter := new(FilterTaking)
    if q.Name != "" {
        filter.Name = q.Name
    } else {
        filter.Name = "%"
    }
    if q.Crew != "" {
        filter.Crew = q.Crew
    } else {
        filter.Crew = "%"
    }
    return filter
}
