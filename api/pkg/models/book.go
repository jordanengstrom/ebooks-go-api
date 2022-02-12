package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Authors       []Author `json:"authors" gorm:"many2many:book_authors"`
	ImgUrl        string   `json:"imgUrl"`
	Title         string   `json:"title"`
	CopyrightYear int      `json:"copyrightYear"`
	About         string   `json:"about"`
}
