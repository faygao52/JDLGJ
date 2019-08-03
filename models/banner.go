package models

import (
	"jdlgj/models/base"

	uuid "github.com/satori/go.uuid"
)

type (
	//Banner describes a banner model
	Banner struct {
		base.Base
		Title    string `json:"title" binding:"required"`
		ImageURI string `json:"imageURI" binding:"required"`
		Link     string `json:"link" binding:"required"`
		Visible  bool   `json:"visible"`
	}

	//BannerResource represents a banner resource for external usage
	BannerResource struct {
		ID       uuid.UUID `json:"id"`
		Title    string    `json:"title"`
		ImageURI string    `json:"imageURI"`
		Link     string    `json:"link"`
		Visible  bool      `json:"visible"`
	}
)

// ToResource Convert banner to external representation
func (b Banner) ToResource() interface{} {
	return BannerResource{
		ID:       b.ID,
		Title:    b.Title,
		ImageURI: b.ImageURI,
		Link:     b.Link,
		Visible:  b.Visible,
	}
}
