package disc

import "cd-catalog-backend-go/structs"

//Disc is a representation of a music disc
type Disc struct {
	structs.Base
	Artist    string `json:"artist" gorm:"not null"`
	Album     string `json:"album"  gorm:"not null"`
	Year      int    `json:"year"  gorm:"not null"`
	Title     string `json:"title"  gorm:"not null;unique;"`
	Style     string `json:"style"  gorm:"not null"`
	SongCount int    `json:"song_count"  gorm:"not null"`
}
