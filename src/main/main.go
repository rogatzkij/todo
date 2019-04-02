package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var db DBtodo

	db.driverName = "mysql"
	db.login = "root"
	db.password = ""
	db.adress = "localhost"
	db.port = 3306
	db.databaseName = "todo"

	err := db.connect()
	if err != nil {
		panic(err)
	}

	//	db.writeUser("test", "pass", "test@test.com")

	usr, ok := db.getUser("test", "pass", "test@test.com")
	if ok {
		fmt.Println(usr.Login)
		fmt.Println(usr.Email)
		fmt.Println(usr.Hash)
	}

	db.WriteCookie("test", "cookie1")
}
