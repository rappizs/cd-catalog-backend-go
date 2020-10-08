package disc

import "cd-catalog-backend-go/common"

//Disc is a representation of a music disc
type Disc struct {
	common.Base
	Artist    string `json:"artist" gorm:"not null" validate:"required,min=3,max=20"`
	Album     string `json:"album"  gorm:"not null" validate:"required,min=3,max=20"`
	Year      int    `json:"year"  gorm:"not null" validate:"required,number,gt=0"`
	Title     string `json:"title"  gorm:"not null;unique;" validate:"required,min=5,max=50"`
	Style     string `json:"style"  gorm:"not null" validate:"required,min=3,max=20"`
	SongCount int    `json:"song_count"  gorm:"not null" validate:"required,number,gt=0"`
}
