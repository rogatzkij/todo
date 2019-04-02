package main

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

type DBtodo struct {
	driverName   string
	login        string
	password     string
	adress       string
	port         int
	databaseName string
	database     *sqlx.DB
}

func (db *DBtodo) connect() error {
	querry := db.login + ":" + "@tcp(" + db.adress + ":" + strconv.Itoa(db.port) + ")/" + db.databaseName

	debug(querry)
	conn, err := sqlx.Connect(db.driverName, querry)
	if err != nil {
		return err
	}

	db.database = conn
	return nil
}
