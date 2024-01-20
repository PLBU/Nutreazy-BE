package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name string
	Gender gender
	YearBorn uint
	Height uint
}

type gender string

const (
	Male gender = "Male"
	Female  gender = "Female"
)