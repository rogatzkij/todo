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
	/*
		var task dbwork.Task
		task.Login = "admin"
		task.Title.Scan("title")
		task.Description.Scan("descr")
		task.Defer.Scan("1")
		task.DateEnd.Scan("12-12-2012")

		ToDoDatabase.AddTask(task)

		task1, b := ToDoDatabase.GetAllTasks("admin")
		fmt.Printf("%v \n err %v", task1, b)
	*/

	serv()
}
