package repository

import (
	"allsounds/pkg/db"
	"fmt"
)

type BaseRepository interface {
	FindAll(offset int, limit int, data interface{}) interface{}
	Search(offset int, limit int, query string, data interface{}) interface{}
}

type Repository struct{}

func (b Repository) FindAll(offset int, limit int, data interface{}, order string) interface{} {
	db.DBCon.Order(fmt.Sprintf("%s asc", order)).Limit(limit).Offset(offset).Find(&data)

	return data
}

func (b Repository) Search(offset int, limit int, query string, data interface{}, order string) interface{} {
	db.DBCon.Order(fmt.Sprintf("%s asc", order)).Limit(limit).Offset(offset).Where(fmt.Sprintf("%s LIKE ?", order), "%"+query+"%").Find(&data)

	return data
}

func (b Repository) FindById(id int, data interface{}) *interface{} {
	db.DBCon.First(&data, id)

	return &data
}
