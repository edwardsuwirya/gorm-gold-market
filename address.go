package main

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Address struct {
	ID         string `gorm:"column:id;size:36;primaryKey"`
	StreetName string `gorm:"size:50;not null"`
	City       string `gorm:"size:10;not null"`
	PostalCode string `gorm:"size:7;not null"`
	CustomerID string `gorm:"column:customer_id;size:36;not null"`
	gorm.Model
}

func (a *Address) TableName() string {
	return "mst_user_address"
}
func (a *Address) ToString() string {
	address, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(address)
}
