package dbwork

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

const (
	msgInfoQuerry = "request generated: %s"
	msgSuccessful = "connection successfully established"
)

// Connect - установление соединения с БД
func (db *DBtodo) Connect() error {

	querry := fmt.Sprintf("%s:@tcp(%s:%d)/%s", db.Login, db.Adress, db.Port, db.DatabaseName)
	db.Log.Infof(msgInfoQuerry, querry)

	conn, err := sqlx.Connect(db.DriverName, querry)
	if err != nil {
		db.Log.Fatal(err.Error())
		return err
	}

	db.Database = conn
	db.Log.Info(msgSuccessful)

	return nil
}

// LogInit - инициализация логгера
func (db *DBtodo) LogInit() {
	db.Log = logrus.New()
	db.Log.Out = os.Stdout
}
