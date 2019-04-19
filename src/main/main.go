package main

import (
	"main/dbwork"

	_ "github.com/go-sql-driver/mysql"
)

var ToDoDatabase dbwork.DBtodo

func main() {
	ToDoDatabase.LogInit()

	err := ToDoDatabase.ParseSQLSettings("SQL.config")
	if err != nil {
		return
	}

	err = ToDoDatabase.Connect()
	if err != nil {
		return
	}

	serv()
}
