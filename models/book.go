package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	ID        uint   `form:"-",gorm:"primary_key"`
	Name      string `form:"name,text,name:"`
	Quantity  int    `form:"quantity,int,quantity:"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
