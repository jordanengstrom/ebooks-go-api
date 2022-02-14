package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	FirstName  string `json:"firstName" validate:"required"`
	MiddleName string `json:"middleName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
}
