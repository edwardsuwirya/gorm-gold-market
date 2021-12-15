package main

import "log"

type CustomerProductRepo struct {
	BaseRepo
}

func (cr *CustomerProductRepo) Insert(newCustomerProduct CustomerProduct) error {
	result := cr.conn.Db.Create(&newCustomerProduct)
	return cr.HandleError(result)
}

func (cr *CustomerProductRepo) FindById(id string) Customer {
	var customer Customer
	result := cr.conn.Db.Debug().Preload("Addresses").First(&customer, "id = ?", id)
	err := cr.HandleError(result)
	if err != nil {
		log.Println(err)
		return Customer{}
	}
	return customer
}

func (cr *CustomerProductRepo) FindAllCustomerPaging(limit int, page int) []Customer {
	var customers []Customer
	result := cr.conn.Db.Limit(limit).Offset((page - 1) * limit).Preload("Addresses").Find(&customers)
	err := cr.HandleError(result)
	if err != nil {
		return nil
	}
	return customers
}

func NewCustomerProductRepo(conn *DbConn) *CustomerProductRepo {
	return &CustomerProductRepo{
		BaseRepo{
			conn: conn,
		},
	}
}
