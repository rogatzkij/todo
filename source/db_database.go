package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// DBtodo - структура для работы с бд
type DBtodo struct {
	DriverName   string         `json:",omitempty"`
	Login        string         `json:",omitempty"`
	Password     string         `json:",omitempty"`
	Adress       string         `json:",omitempty"`
	Port         int            `json:",omitempty"`
	DatabaseName string         `json:",omitempty"`
	Database     *sqlx.DB       `json:"-"`
	Log          *logrus.Logger `json:"-"`
}

// Connect - установление соединения с БД
func (db *DBtodo) Connect() error {
	querry := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.Login, db.Password, db.Adress, db.Port, db.DatabaseName)
	db.Log.Infof("Request generated: %s", querry)

	conn, err := sqlx.Connect(db.DriverName, querry)
	if err != nil {
		db.Log.Fatal(err.Error())
		return err
	}

	db.Database = conn
	db.Log.Info("Connection successfully established")

	return nil
}

// LogInit - инициализация логгера
func (db *DBtodo) LogInit() {
	db.Log = logrus.New()
	db.Log.Out = os.Stdout
}