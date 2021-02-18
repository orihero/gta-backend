package models

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	Firstname string
	Lastname string
	DateOfBirth time.Time
	Salary float64
}
