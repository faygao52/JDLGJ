package models

import (
	"github.com/jinzhu/gorm"
)

type (
	//ServiceCenter describes a serivce center model type
	ServiceCenter struct {
		gorm.Model
		Name        string `json:"name"`
		WorkingHour string `json:"workingHour"`
		Contact     string `json:"contact"`
		Address     string `json:"address"`
	}

	//ServiceCenterResource represents service center for external usage
	ServiceCenterResource struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		WorkingHour string `json:"workingHour"`
		Contact     string `json:"contact"`
		Address     string `json:"address"`
	}
)
