package models

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name     string `json:"name"`
	Location string `json:"location"`
	Price    int    `json:"price"`
}
