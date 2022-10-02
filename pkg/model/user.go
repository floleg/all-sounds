package model

type User struct {
	Entity
	Login  string  `faker:"username"`
	Tracks []Track `gorm:"many2many:user_tracks;" faker:"-"`
}
