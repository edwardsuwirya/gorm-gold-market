package main

import (
	"encoding/json"
	"gorm.io/gorm"
)

type UserCredential struct {
	ID         string `gorm:"column:id;size:36;primaryKey"`
	Username   string `gorm:"size:50;not null"`
	Password   string `gorm:"size:10;not null"`
	Email      string `gorm:"size:50;not null"`
	CustomerID string
	IsActive   bool
	gorm.Model
}

func (c *UserCredential) TableName() string {
	return "mst_user_credential"
}
func (c *UserCredential) ToString() string {
	cred, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(cred)
}
