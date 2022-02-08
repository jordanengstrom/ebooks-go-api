package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}
