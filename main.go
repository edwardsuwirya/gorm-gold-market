package main

import "log"

func main() {
	db := NewDbConn()
	defer db.Close()
	db.PingTest()
	db.Migration(&Customer{}, &UserCredential{})

	customerRepo := NewCustomerRepo(db)
	//err := customerRepo.Insert(Customer{
	//	ID:        "C003",
	//	FirstName: "Tika",
	//	LastName:  "Yesi",
	//	BirthDate: time.Time{},
	//	Address:   "Lampung",
	//	Status:    1,
	//	UserCredential: UserCredential{
	//		ID:       "C003",
	//		Username: "tika",
	//		Password: "222333",
	//		Email:    "tika@enigmacamp.com",
	//		IsActive: true,
	//	},
	//})
	//if err != nil {
	//	log.Println(err.Error())
	//}

	// Delete menggunakan ID, coba selain menggunakan ID, apa yang terjadi
	//err = customerRepo.Delete(Customer{ID: "C001"})
	//if err != nil {
	//	log.Println(err.Error())
	//}

	//Update
	//err = customerRepo.Update(Customer{
	//	ID:        "C001",
	//	FirstName: "Kaka",
	//	LastName:  "Reza",
	//	BirthDate: time.Time{},
	//	Address:   "Depok 1",
	//	Status:    0,
	//	Username:  "",
	//	Password:  "",
	//	Email:     "",
	//})

	authRepo := NewAuthRepo(db)
	isAuth := authRepo.Authenticate(UserCredential{
		Username: "tika",
		Password: "222333",
	})
	log.Println(isAuth)
	customer := customerRepo.FindById("C003")
	log.Println(customer.ToString())

	//customers := customerRepo.FindAllCustomerPaging(2, 1)
	//log.Println("Page 1")
	//for _, c := range customers {
	//	log.Println(c.ToString())
	//}
	//customers = customerRepo.FindAllCustomerPaging(2, 2)
	//log.Println("Page 2")
	//for _, c := range customers {
	//	log.Println(c.ToString())
	//}

	//customers := customerRepo.FindCustomerPaging(Customer{FirstName: "Ka"}, 2, 1)
	//log.Println("Page 1")
	//for _, c := range customers {
	//	log.Println(c.ToString())
	//}
	//
	//total := customerRepo.TotalCustomer()
	//log.Println(total)
}
