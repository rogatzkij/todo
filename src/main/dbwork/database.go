package dbwork

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

type DBtodo struct {
	DriverName   string
	Login        string
	Password     string
	Adress       string
	Port         int
	DatabaseName string
	Database     *sqlx.DB
}

func (db *DBtodo) Connect() error {
	querry := db.Login + ":" + "@tcp(" + db.Adress + ":" + strconv.Itoa(db.Port) + ")/" + db.DatabaseName

	//debug(querry)
	conn, err := sqlx.Connect(db.DriverName, querry)
	if err != nil {
		return err
	}

	db.Database = conn
	return nil
}
