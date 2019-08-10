package repository

import (
	"fmt"
	"jdlgj/models/base"
	"jdlgj/repository/db"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
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

//FindBy given a field find and return the first element
func FindBy(field string, value base.ModelInterface, query string) (interface{}, error) {
	err := db.GetDB().Where(field+" = ?", query).First(value).Error
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
func List(collection interface{}, page string, size string, orderBy []string) *pagination.Paginator {
	pageInt, _ := strconv.Atoi(page)
	limit, _ := strconv.Atoi(size)
	return pagination.Paging(&pagination.Param{
		DB:      db.GetDB(),
		Page:    pageInt,
		Limit:   limit,
		OrderBy: orderBy,
		ShowSQL: true,
	}, collection)
}

//SearchAll Search for results and return pagination result
func SearchAll(collection interface{}, page string, size string, orderBy []string, query string, fields []string) *pagination.Paginator {
	db := db.GetDB()
	fmt.Printf("query: %s", query)

	for i, field := range fields {
		if i == 0 {
			db = db.Where(field+" LIKE ?", "%"+query+"%")
		} else {
			db = db.Or(field+" LIKE ?", "%"+query+"%")
		}
	}
	pageInt, _ := strconv.Atoi(page)
	limit, _ := strconv.Atoi(size)
	return pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    pageInt,
		Limit:   limit,
		OrderBy: orderBy,
		ShowSQL: false,
	}, collection)
}
