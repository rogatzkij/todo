package main

import (
	_ "github.com/go-sql-driver/mysql"
)

var ToDoDatabase DBtodo

func main() {
	// запускаем логгер
	ToDoDatabase.LogInit()

	// читаем конфиги по коннекту с БД
	err := ToDoDatabase.ParseSQLSettings("SQL.config")
	if err != nil {
		return
	}

	// коннектимся с БД
	err = ToDoDatabase.Connect()
	if err != nil {
		return
	}

	// слушаем киентов
	serv()
}
