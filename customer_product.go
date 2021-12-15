package main

import (
	"encoding/json"
	"gorm.io/gorm"
)

// Customer Customers : Untuk konfigurasi field tags, bisa dilihat di
//https://gorm.io/docs/models.html#Fields-Tags
type CustomerProduct struct {
	ID          string      `gorm:"column:id;size:36;primaryKey"`
	ProductName string      `gorm:"column:first_name;size:50;not null"`
	Customers   []*Customer `gorm:"many2many:customer_products;"`
	gorm.Model
}

func (c *CustomerProduct) TableName() string {
	return "mst_customer_product"
}
func (c *CustomerProduct) ToString() string {
	customer, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(customer)
}
