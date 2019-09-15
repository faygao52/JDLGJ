package models

import (
	"jdlgj/models/base"
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	//User backend user in this platform
	User struct {
		base.Base
		Username string `json:"username" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}

	//WcUser wechat user using wechat miniapp
	WcUser struct {
		base.Base
		OpenID string `json:"openId" gorm:"unique_index"`
	}

	//UserResource represents user for external usage
	UserResource struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"name" binding:"required"`
		Role      string    `json:"role" binding:"required"`
		CreatedAt time.Time `json:"createdAt"  time_format:"2006-01-02"`
		UpdatedAt time.Time `json:"updatedAt"  time_format:"2006-01-02"`
	}

	//WcUserResource represents wc user
	WcUserResource struct {
		ID        uuid.UUID `json:"id"`
		OpenID    string    `json:"openId" gorm:"unique_index"`
		CreatedAt time.Time `json:"createdAt"  time_format:"2006-01-02"`
	}
)

// ToResource converts a User into external representation
func (us User) ToResource() interface{} {
	return UserResource{
		ID:        us.ID,
		Username:  us.Username,
		Role:      us.Role,
		CreatedAt: us.CreatedAt,
		UpdatedAt: us.UpdatedAt,
	}
}

// ToResource converts a User into external representation
func (wc WcUser) ToResource() interface{} {
	return WcUserResource{
		ID:        wc.ID,
		OpenID:    wc.OpenID,
		CreatedAt: wc.CreatedAt,
	}
}
