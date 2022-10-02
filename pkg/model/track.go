package model

type Track struct {
	Entity
	Title    string `faker:"sentence"`
	ArtistID uint
	Users    []User  `gorm:"many2many:user_tracks;" faker:"-"`
	Albums   []Album `gorm:"many2many:album_tracks;" faker:"-"`
}
