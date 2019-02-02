package models

import (
	"strings"
	"time"
)

// FetchQuery the basic fetch query
type FetchQuery struct {
	Limit  int         `form:"limit"`
	Page   int         `form:"page"`
	Order  string      `form:"order"`
	Dates  []time.Time `form:"dates"`
	Fields string      `form:"fields"`
}

func (q *FetchQuery) sanitize() {
	if q.Limit == 0 {
		q.Limit = 10
	}
	if q.Page == 0 {
		q.Page = 1
	}
	if q.Order == "" {
		q.Order = "timestamp desc"
	}
	if len(q.Dates) == 0 {
		q.Dates = []time.Time{}
	}
	if q.Fields == "" {
		q.Fields = "*"
	} else if !strings.Contains(q.Fields, "timestamp") {
		q.Fields += ",timestamp"
	}
}
