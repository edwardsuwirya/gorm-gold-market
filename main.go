package main

func main() {
	db := NewDbConn()
	defer db.Close()
	db.PingTest()
	db.Migration(&Customer{}, &UserCredential{}, &Address{}, &CustomerProduct{})

	customerProductRepo := NewCustomerProductRepo(db)
	customerProductRepo.FindTotalCustomerByProduct()
	//customerProductRepo.Insert(CustomerProduct{
	//	ID:          "CP001",
	//	ProductName: "Deposito Rupiah",
	//})
	//customerProductRepo.Insert(CustomerProduct{
	//	ID:          "CP002",
	//	ProductName: "Deposito Dollar",
	//})
	//customerProductRepo.Insert(CustomerProduct{
	//	ID:          "CP003",
	//	ProductName: "Tabungan Emas",
	//})
	//customerProductRepo.Insert(CustomerProduct{
	//	ID:          "CP004",
	//	ProductName: "Tabungan",
	//})
	//customerRepo := NewCustomerRepo(db)
	//err := customerRepo.OpenProductForExistingCustomer(Customer{
	//	ID: "C001",
	//	Products: []*CustomerProduct{
	//		{
	//			ID: "CP003",
	//		},
	//	},
	//})
	//if err != nil {
	//	return
	//}
	//err := customerRepo.Insert(Customer{
	//	ID:        "C001",
	//	FirstName: "Jution",
	//	LastName:  "Chandra",
	//	BirthDate: time.Time{},
	//	Addresses: []Address{
	//		{
	//			ID:         "A003",
	//			StreetName: "Jl.Jambu",
	//			City:       "Lampung",
	//			PostalCode: "111222",
	//		},
	//		{
	//			ID:         "A004",
	//			StreetName: "Jl.Mangga",
	//			City:       "Lampung",
	//			PostalCode: "111222",
	//		},
	//	},
	//	Status: 1,
	//	UserCredential: UserCredential{
	//		ID:       "C001",
	//		Username: "jution",
	//		Password: "111222",
	//		Email:    "jution@enigmacamp.com",
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

	//authRepo := NewAuthRepo(db)
	//isAuth := authRepo.Authenticate(UserCredential{
	//	Username: "tika",
	//	Password: "222333",
	//})
	//log.Println(isAuth)
	//customer := customerRepo.FindById("C001")
	//log.Println(customer.ToString())

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
