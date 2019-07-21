package models

import (
	"github.com/jinzhu/gorm"
)

type (
	//LawFirm describes a law firm model
	LawFirm struct {
		gorm.Model
		Title       string  `json:"title"`
		WorkingHour string  `json:"workingHour"`
		Contact     string  `json:"contact"`
		Address     string  `json:"address"`
		Description string  `json:"description"`
		Review      float32 `json:"review"`
		Services    int     `jons:"service"`
	}

	//LawFirmResource representa law firm resource for external usage
	LawFirmResource struct {
		ID          uint    `json:"id"`
		Title       string  `json:"title"`
		WorkingHour string  `json:"workingHour"`
		Contact     string  `json:"contact"`
		Address     string  `json:"address"`
		Description string  `json:"description"`
		Review      float32 `json:"review"`
		Services    int     `jons:"service"`
	}
)
