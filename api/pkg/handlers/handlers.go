package handlers

import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

func New(initializedDb *gorm.DB) handler {
	return handler{initializedDb}
}
