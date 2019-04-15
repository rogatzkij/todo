package main

import (
	"main/dbwork"

	_ "github.com/go-sql-driver/mysql"
)

var ToDoDatabase dbwork.DBtodo

func main() {

	ToDoDatabase.DriverName = "mysql"
	ToDoDatabase.Login = "root"
	ToDoDatabase.Password = ""
	ToDoDatabase.Adress = "localhost"
	ToDoDatabase.Port = 3306
	ToDoDatabase.DatabaseName = "todo"

	err := ToDoDatabase.Connect()
	if err != nil {
		panic(err)
	}

	serv()
}
