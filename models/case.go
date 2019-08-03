package models

import (
	"jdlgj/models/base"
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	//Case describes a case model
	Case struct {
		base.Base
		Catalog  string `json:"catalog"  binding:"required"`
		Question string `json:"question"  binding:"required"`
		Answer   string `json:"answer"  binding:"required"`
		Contact  string `json:"contact"`
		Lawyer   string `json:"lawyer"`
	}

	//CaseResource represents a case resource for external usage
	CaseResource struct {
		ID        uuid.UUID `json:"id"`
		Catalog   string    `json:"catalog"`
		Question  string    `json:"question"`
		Answer    string    `json:"answer"`
		Contact   string    `json:"contact"`
		Lawyer    string    `json:"lawyer"`
		CreatedAt time.Time `json:"createdAt"  time_format:"2006-01-02"`
	}
)

// ToResource Convert case to external representation
func (c Case) ToResource() interface{} {
	return CaseResource{
		ID:        c.ID,
		Catalog:   c.Catalog,
		Question:  c.Question,
		Answer:    c.Answer,
		Contact:   c.Contact,
		Lawyer:    c.Lawyer,
		CreatedAt: c.CreatedAt,
	}
}
