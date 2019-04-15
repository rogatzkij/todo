package dbwork

import (
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DBtodo struct {
	DriverName   string   `json:",omitempty"`
	Login        string   `json:",omitempty"`
	Password     string   `json:",omitempty"`
	Adress       string   `json:",omitempty"`
	Port         int      `json:",omitempty"`
	DatabaseName string   `json:",omitempty"`
	Database     *sqlx.DB `json:"-"`
	Log          *logrus.Logger
}

func (db *DBtodo) Connect() error {
	querry := db.Login + ":" + "@tcp(" + db.Adress + ":" + strconv.Itoa(db.Port) + ")/" + db.DatabaseName

	db.Log.Info("Запрос: " + querry)

	conn, err := sqlx.Connect(db.DriverName, querry)
	if err != nil {
		db.Log.Fatal(err.Error())
		return err
	}

	db.Database = conn
	db.Log.Info("Соединение успешно установленно")

	return nil
}

func (db *DBtodo) LogInit() {

	db.Log = logrus.New()
	db.Log.Out = os.Stdout
}
