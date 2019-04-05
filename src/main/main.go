package main

import (
	_ "github.com/go-sql-driver/mysql"
)

var ToDoDatabase DBtodo

func main() {

	ToDoDatabase.driverName = "mysql"
	ToDoDatabase.login = "root"
	ToDoDatabase.password = ""
	ToDoDatabase.adress = "localhost"
	ToDoDatabase.port = 3306
	ToDoDatabase.databaseName = "todo"

	err := ToDoDatabase.connect()
	if err != nil {
		panic(err)
	}

	serv()
}
