package base

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//ModelInterface a model with a toResource
type ModelInterface interface {
	ToResource() interface{}
}

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// PaginationResource is external representor for paginated resource
type PaginationResource struct {
	DataCollection interface{} `json:"dataCollection"`
	TotalElement   int         `json:"totalElement"`
	CurrentPage    int         `json:"currentPage"`
	TotalPages     int         `json:"totalPages"`
}
