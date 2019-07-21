package models

import (
	"github.com/jinzhu/gorm"
)

type (
	//Banner describes a banner model
	Banner struct {
		gorm.Model
		Title    string `json:"title" binding:"required"`
		ImageURI string `json:"imageURI" binding:"required"`
		Link     string `json:"link" binding:"required"`
		Visible  bool   `json:"visible"`
	}

	//BannerResource represents a banner resource for external usage
	BannerResource struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		ImageURI string `json:"imageURI"`
		Link     string `json:"link"`
		Visible  bool   `json:"visible"`
	}
)
