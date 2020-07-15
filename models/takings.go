package models

import "strconv"

type (
	TakingOld struct {
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
		Received int      `json:"received" validate:"required"`
		User     []User   `json:"involvedSupporter" validate:"required"`
		Sources  []Source `json:"sources" validate:"required"`
	}

	User struct {
		Uuid string `json:"uuid" validate:"required"`
		Name string `json:"name" validate:"required"`
		Tag  string `json:"tag"`
	}
	UserList []User

	SourceCreate struct {
		Category     string       `json:"category" validate:"required"`
		Amount       int          `json:"amount" validate:"required"`
		Currency     string       `json:"currency" validate:"required"`
		Description  string       `json:"description"`
		TypeOfSource TypeOfSource `json:"typeOfSource" validate:"required"`
		Norms        string       `json:"norms" validate:"required"`
	}
	Source struct {
		Uuid         string `json:"uuid" validate:"required"`
		Category     string `json:"category" validate:"required"`
		Amount       int    `json:"amount" validate:"required"`
		Currency     string `json:"currency" validate:"required"`
		Description  string `json:"description"`
		TypeOfSource string `json:"typeOfSource" validate:"required"`
		Norms        string `json:"norms" validate:"required"`
	}
	SourceList []Source

	TypeOfSource struct {
		Category string   `json:"category" validate:"required"`
		External External `json:"external"`
	}

	Confirmed struct {
		Date int  `json:"date" validate:"required"`
		User User `json:"user" validate:"required"`
	}

	External struct {
		Location      string `json:"location" validate:"required"`
		ContactPerson string `json:"contactPerson" validate:"required"`
		Email         string `json:"email" validate:"required"`
		Address       string `json:"address" validate:"required"`
		Receipt       bool   `json:"receipt" validate:"required"`
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
	CrewList []Crew

	Money struct {
		Amount   int    `json:"amount" validate:"required"`
		Currency string `json:"currency" validate:"required"`
		Status   string `json:"status" validate:"required"`
	}
	MoneyList []Money
	Taking    struct {
		Uuid        string     `json:"uuid" validate:"required"`
		Received    int        `json:"received" validate:"required"`
		Description string     `json:"description" validate:"required"`
		Comment     string     `json:"comment" validate:"required"`
		Category    string     `json:"birthdate" validate:"required"`
		Author      User       `json:"author"`
		User        UserList   `json:"supporter"`
		Crews       CrewList   `json:"crews"`
		Sources     SourceList `json:"sources"`
		Money       []Money    `json:"Money"`
	}

	// TODO: Finish struct
	TakingCreate struct {
		Received         int            `json:"received" validate:"required"`
		Description      string         `json:"description" validate:"required"`
		Comment          CommentCreate  `json:"comment" validate:"required"`
		Category         string         `json:"category" validate:"required"`
		ReasonForPayment string         `json:"reasonForPayment"`
		Author           User           `json:"author"`
		User             UserList       `json:"supporter"`
		Crews            CrewList       `json:"crews"`
		Sources          []SourceCreate `json:"sources"`
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

func (list *SourceList) Distinct() *SourceList {
	r := make(SourceList, 0, len(*list))
	m := make(map[Source]bool)
	for _, val := range *list {
		if _, ok := m[val]; !ok {
			m[val] = true
			r = append(r, val)
		}
	}
	return &r
}
func (list *UserList) Distinct() *UserList {
	r := make(UserList, 0, len(*list))
	m := make(map[User]bool)
	for _, val := range *list {
		if _, ok := m[val]; !ok {
			m[val] = true
			r = append(r, val)
		}
	}
	return &r
}
func (list *MoneyList) Distinct() *MoneyList {
	r := make(MoneyList, 0, len(*list))
	m := make(map[Money]bool)
	for _, val := range *list {
		if _, ok := m[val]; !ok {
			m[val] = true
			r = append(r, val)
		}
	}
	return &r
}
func (list *CrewList) Distinct() *CrewList {
	r := make(CrewList, 0, len(*list))
	m := make(map[Crew]bool)
	for _, val := range *list {
		if _, ok := m[val]; !ok {
			m[val] = true
			r = append(r, val)
		}
	}
	return &r
}
