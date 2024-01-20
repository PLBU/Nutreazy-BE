package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name string
	Gender Gender
	YearBorn uint
	Height uint
}

type Gender string

const (
	Male Gender = "Male"
	Female  Gender = "Female"
)