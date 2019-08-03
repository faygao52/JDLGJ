package models

import (
	"jdlgj/models/base"

	uuid "github.com/satori/go.uuid"
)

type (
	//ServiceCenter describes a serivce center model type
	ServiceCenter struct {
		base.Base
		Name        string  `json:"name" binding:"required"`
		WorkingHour string  `json:"workingHour" binding:"required"`
		Contact     string  `json:"contact" binding:"required"`
		Address     string  `json:"address" binding:"required"`
		Longitude   float64 `json:"longitude" binding:"required"`
		Latitude    float64 `json:"latitude" binding:"required"`
	}

	//ServiceCenterResource represents service center for external usage
	ServiceCenterResource struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		WorkingHour string    `json:"workingHour"`
		Contact     string    `json:"contact"`
		Address     string    `json:"address"`
		Longitude   float64   `json:"longitude"`
		Latitude    float64   `json:"latitude"`
	}
)

// ToResource converts a ServiceCenter into external representation
func (sc ServiceCenter) ToResource() interface{} {
	return ServiceCenterResource{
		ID:          sc.ID,
		Name:        sc.Name,
		WorkingHour: sc.WorkingHour,
		Address:     sc.Address,
		Contact:     sc.Contact,
		Longitude:	 sc.Longitude,
		Latitude:	 sc.Latitude,
	}
}
