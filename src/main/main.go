package main

import (
	"main/dbwork"

	_ "github.com/go-sql-driver/mysql"
)

var ToDoDatabase dbwork.DBtodo

func main() {
	ToDoDatabase.LogInit()
	ToDoDatabase.ParseSQLSettings("SQL.config")

	err := ToDoDatabase.Connect()
	if err != nil {
		return
	}

	serv()
}
