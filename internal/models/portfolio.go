package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model
	Name   string
	Assets []Asset
}
