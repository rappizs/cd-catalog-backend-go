package structs

import uuid "github.com/satori/go.uuid"

//Disc is a representation of a music disc
type Disc struct {
	Base
	Artist    string `json:"artist" gorm:"not null"`
	Album     string `json:"album"  gorm:"not null"`
	Year      int    `json:"year"  gorm:"not null"`
	Title     string `json:"title"  gorm:"not null;unique;"`
	Style     string `json:"style"  gorm:"not null"`
	SongCount int    `json:"song_count"  gorm:"not null"`
}

//Artist is a representation of a Disc's artist
type Artist struct {
	Base
	Name uint `json:"name" gorm:"unique"`
}

//Style is a representation of a music style
type Style struct {
	Base
	Name uint `json:"name" gorm:"unique"`
}

//Base contains common columns for all structs
type Base struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
}
