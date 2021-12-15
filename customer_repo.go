package main

type CustomerRepo struct {
	BaseRepo
}

func (cr *CustomerRepo) Insert(newCustomer Customer) error {
	result := cr.conn.Db.Create(&newCustomer)
	return cr.HandleError(result)
}

func (cr *CustomerRepo) Delete(customer Customer) error {
	// Secara default akan melakukan soft delete
	result := cr.conn.Db.Delete(&customer)

	//Hard Delete
	//result := cr.conn.Db.Unscoped().Delete(&customer)

	return cr.HandleError(result)
}

func (cr *CustomerRepo) Update(updateCustomerInfo Customer) error {
	//Kalau tidak ada primary key nya diberikan akan melakukan insert
	//result := cr.conn.Db.Save(&updateCustomerInfo)

	//Ketika menggunakan struct updates, gorm hanya melakukan update yang bukan zero value
	//result := cr.conn.Db.Model(&updateCustomerInfo).Updates(updateCustomerInfo)

	//Cara melakukan dengan zero value
	result := cr.conn.Db.Model(&updateCustomerInfo).Select("*").Updates(updateCustomerInfo)
	return cr.HandleError(result)
}

func (cr *CustomerRepo) FindById(id string) Customer {
	var customer Customer
	result := cr.conn.Db.First(&customer, "id = ?", id)
	err := cr.HandleError(result)
	if err != nil {
		return Customer{}
	}
	return customer
}

func (cr *CustomerRepo) FindAllCustomerPaging(limit int, page int) []Customer {
	var customers []Customer
	result := cr.conn.Db.Limit(limit).Offset((page - 1) * limit).Find(&customers)
	err := cr.HandleError(result)
	if err != nil {
		return nil
	}
	return customers
}

func (cr *CustomerRepo) FindCustomerPaging(customer Customer, limit int, page int) []Customer {
	var customers []Customer
	result := cr.conn.Db.Limit(limit).Offset((page - 1) * limit).Where(&customer).Find(&customers)
	err := cr.HandleError(result)
	if err != nil {
		return nil
	}
	return customers
}

func (cr *CustomerRepo) TotalCustomer() AggResult {
	var total AggResult
	result := cr.conn.Db.Model(&Customer{}).Select("count(id) as total").Find(&total)
	err := cr.HandleError(result)
	if err != nil {
		return AggResult{
			Total: -1,
		}
	}
	return total
}

func NewCustomerRepo(conn *DbConn) *CustomerRepo {
	return &CustomerRepo{
		BaseRepo{
			conn: conn,
		},
	}
}
