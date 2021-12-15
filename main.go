package main

import (
	"log"
)

func main() {
	db := NewDbConn()
	defer db.Close()
	db.PingTest()
	db.Migration(&Customer{})

	customerRepo := NewCustomerRepo(db)
	//err := customerRepo.Insert(Customer{
	//	ID:        "C003",
	//	FirstName: "Tika",
	//	LastName:  "Yesi",
	//	BirthDate: time.Time{},
	//	Address:   "Lampung",
	//	Status:    1,
	//	Username:  "tika",
	//	Password:  "222333",
	//	Email:     "tika@enigmacamp.com",
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

	//customer := customerRepo.FindById("C001")
	//log.Println(customer.ToString())
	//
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

	customers := customerRepo.FindCustomerPaging(Customer{FirstName: "Ka"}, 2, 1)
	log.Println("Page 1")
	for _, c := range customers {
		log.Println(c.ToString())
	}

	total := customerRepo.TotalCustomer()
	log.Println(total)
}
