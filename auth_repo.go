package main

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type AuthRepo struct {
	BaseRepo
}

func (cr *AuthRepo) Authenticate(cred UserCredential) error {
	//Alternatif lain menggunakan WHERE
	var userCred UserCredential
	result := cr.conn.Db.Where("username = ? AND password = ?", cred.Username, cred.Password).First(&userCred)
	//result := cr.conn.Db.Preload("mst_user_credential").Where("username = ? AND password = ?", cred.Username, cred.Password).First(&userCred)
	err := cr.HandleError(result)
	if err != nil {
		//GORM returns ErrRecordNotFound when failed to find data with First, Last, Take
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Unauthorized")
		}
		return err
	}
	return nil
}

func (cr *AuthRepo) Authenticate2(cred UserCredential) (Customer, error) {
	var userCred UserCredential
	result := cr.conn.Db.Where("username = ? AND password = ?", cred.Username, cred.Password).Preload("mst_customers").First(&userCred)
	err := cr.HandleError(result)
	log.Println(result)
	if err != nil {
		//GORM returns ErrRecordNotFound when failed to find data with First, Last, Take
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Customer{}, errors.New("Unauthorized")
		}
		return Customer{}, err
	}
	return Customer{}, nil
}

func NewAuthRepo(conn *DbConn) *AuthRepo {
	return &AuthRepo{
		BaseRepo{
			conn: conn,
		},
	}
}
