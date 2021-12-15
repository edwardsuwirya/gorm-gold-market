package main

import "gorm.io/gorm"

type BaseRepo struct {
	conn *DbConn
}

func (br *BaseRepo) HandleError(db *gorm.DB) error {
	if db.Error != nil {
		return db.Error
	}
	return nil
}
