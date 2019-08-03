package repository

import (
	"jdlgj/models/base"
	"jdlgj/repository/db"

	p "github.com/Prabandham/paginator"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//Create a new record if it is not exist
func Create(value base.ModelInterface) interface{} {
	db.GetDB().Create(value)
	return value.ToResource()
}

//FindByID returns a record by its uuid
func FindByID(value base.ModelInterface, id uuid.UUID) (interface{}, error) {
	err := db.GetDB().Where("id = ?", id).First(value).Error
	return value.ToResource(), err
}

//Update a existing record
func Update(c *gin.Context, value base.ModelInterface, id uuid.UUID) (interface{}, error) {
	if _, err := FindByID(value, id); err != nil {
		return nil, err
	}
	c.BindJSON(value)
	db.GetDB().Save(value)
	return value.ToResource(), nil
}

//DeleteByID soft remove a record by its uuid
func DeleteByID(value base.ModelInterface, id uuid.UUID) error {
	if _, err := FindByID(value, id); err != nil {
		return err
	}
	db.GetDB().Delete(value)
	return nil
}

//List all records with pagination
func List(collection interface{}, page string, size string, orderBy []string) *p.Data {
	paginator := p.Paginator{
		DB:      db.GetDB(),
		Page:    page,
		PerPage: size,
		OrderBy: orderBy,
	}
	return paginator.Paginate(collection)
}
