package model

// User data entity
type User struct {
	Entity
	Login  string  `faker:"username"`
	Tracks []Track `gorm:"many2many:user_tracks;" faker:"-"`
}
