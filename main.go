package main

func main() {
	db := NewDbConn()
	defer db.Close()
	db.PingTest()
	db.Migration(&Customer{})
}
