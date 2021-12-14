package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

// Customers : Untuk konfigurasi field tags, bisa dilihat di
//https://gorm.io/docs/models.html#Fields-Tags
type Customers struct {
	ID        string `gorm:"column:id;size:36;primaryKey"`
	FirstName string `gorm:"column:first_name;size:50;not null"`
	LastName  string `gorm:"column:last_name;size:50;not null"`
	BirthDate time.Time
	Address   string
	Status    int
	Username  string
	Password  string
	Email     string
	gorm.Model
}

func (c *Customers) TableName() string {
	return "mst_customers"
}
func (c *Customers) ToString() string {
	customer, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(customer)
}
