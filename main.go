package main

func main() {
	db := NewDbConn()
	defer db.Close()
	db.Migration(&Customers{})
}
