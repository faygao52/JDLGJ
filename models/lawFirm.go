package models

import (
	"jdlgj/models/base"

	uuid "github.com/satori/go.uuid"
)

type (
	//LawFirm describes a law firm model
	LawFirm struct {
		base.Base
		Title       string  `json:"title"  binding:"required"`
		WorkingHour string  `json:"workingHour"`
		Contact     string  `json:"contact"`
		Address     string  `json:"address"  binding:"required"`
		Description string  `json:"description"`
		Reviews     float32 `json:"review"`
		Services    int     `jons:"services"`
		Icon        string  `json:"icon"`
	}

	//LawFirmResource representa law firm resource for external usage
	LawFirmResource struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		WorkingHour string    `json:"workingHour"`
		Contact     string    `json:"contact"`
		Address     string    `json:"address"`
		Description string    `json:"description"`
		Reviews     float32   `json:"review"`
		Services    int       `jons:"services"`
		Icon        string    `json:"icon"`
	}
)

// ToResource Convert LawFirm to external representation
func (lf *LawFirm) ToResource() interface{} {
	return LawFirmResource{
		ID:          lf.ID,
		Title:       lf.Title,
		WorkingHour: lf.WorkingHour,
		Contact:     lf.Contact,
		Address:     lf.Address,
		Description: lf.Description,
		Reviews:     lf.Reviews,
		Services:    lf.Services,
		Icon:        lf.Icon,
	}
}
