package repository

import (
	"allsounds/pkg/db"
)

type BaseRepository interface {
	FindAll(offset int, limit int, data interface{}) interface{}
	Search(offset int, limit int, query string, data interface{}) interface{}
}

type repository struct{}

func (b repository) FindAll(offset int, limit int, data interface{}) interface{} {
	db.DBCon.Order("title asc").Limit(limit).Offset(offset).Find(&data)

	return data
}

func (b repository) Search(offset int, limit int, query string, data interface{}) interface{} {
	db.DBCon.Order("title asc").Limit(limit).Offset(offset).Where("title LIKE ?", "%"+query+"%").Find(&data)

	return data
}

func (b repository) FindById(id int, data interface{}) interface{} {
	db.DBCon.First(&data, id)

	return data
}
