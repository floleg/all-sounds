package repository

import (
	"allsounds/pkg/db"
	"fmt"
)

type BaseRepository interface {
	FindAll(offset int, limit int, data interface{}) interface{}
	Search(offset int, limit int, query string, data interface{}) interface{}
}

// Repository exposes generic gorm persistence methods:
//   - FindAll
//   - Search
//   - FindById
type Repository struct{}

// FindAll performs a select query and returns an interface of the given data parameter, filtered with:
//   - offset
//   - limit
//
// The results will be ordered based on the column name passed in the order parameter.
func (b Repository) FindAll(offset int, limit int, data interface{}, order string) interface{} {
	db.DBCon.Order(fmt.Sprintf("%s asc", order)).Limit(limit).Offset(offset).Find(&data)

	return data
}

// Search performs a select query and returns an interface of the given data parameter, filtered with:
//   - offset
//   - limit
//   - query: where clause
//
// The results will be ordered based on the column name passed in the order parameter.
func (b Repository) Search(offset int, limit int, query string, data interface{}, order string) interface{} {
	db.DBCon.Order(fmt.Sprintf("%s asc", order)).Limit(limit).Offset(offset).
		Where(fmt.Sprintf("%s LIKE ?", order), "%"+query+"%").Find(&data)

	return data
}

// FindById retrieves an interface by id, based on the implementation of the input data parameter
func (b Repository) FindById(id int, data interface{}) *interface{} {
	db.DBCon.First(&data, id)

	return &data
}
