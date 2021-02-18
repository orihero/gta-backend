package models

import (
	"gorm.io/gorm"
)

type CarModel struct {
	gorm.Model
	BrandName string
	Year string
	CarID uint
}

