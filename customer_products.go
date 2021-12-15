package main

type CustomerProducts struct {
	CustomerProductID string `gorm:"primaryKey"`
	CustomerID        string `gorm:"primaryKey"`
}
