package models

import "gorm.io/datatypes"

type Author struct {
	ID        uint
	FirstName string
	LastName  string
	Birth     datatypes.Date
}
