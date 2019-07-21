package models

import (
	"github.com/jinzhu/gorm"
)

type (
	//Case describes a case model
	Case struct {
		gorm.Model
		Catalog     string `json:"catalog"`
		Question    string `json:"question"`
		Answer      string `json:"answer"`
		Description string `json:"description"`
		Contact     string `json:"contact"`
		Lawyer      string `json:"lawyer"`
	}

	//CaseResource represents a case resource for external usage
	CaseResource struct {
		ID          uint   `json:"id"`
		Catalog     string `json:"catalog"`
		Question    string `json:"question"`
		Answer      string `json:"answer"`
		Description string `json:"description"`
		Contact     string `json:"contact"`
		Lawyer      string `json:"lawyer"`
	}
)
