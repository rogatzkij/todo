package dbwork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	msgErrorOpenFile   = "could not open file %s"
	msgErrorCreateFile = "failed to create file %s"
	msgErrorBrokenFile = "file %s is broken"
	msgInfoCreateFile  = "created file %s with default settings"
)

// ParseSQLSettings - чтение настроек из конфиг файла
func (todo *DBtodo) ParseSQLSettings(path string) error {

	bytes, err := ioutil.ReadFile(path)
	if err != nil { // если файла нет создаем с настройками поумолчанию
		todo.Log.Errorf(msgErrorOpenFile, path)

		todo.DriverName = "mysql"
		todo.Login = "root"
		todo.Password = ""
		todo.Adress = "localhost"
		todo.Port = 3306
		todo.DatabaseName = "todo"

		bytes, _ := json.Marshal(todo)

		err = ioutil.WriteFile(path, bytes, 0644)
		if err != nil {
			todo.Log.Errorf(msgErrorCreateFile, path)
		}

		todo.Log.Infof(msgInfoCreateFile, path)
		return nil
	}

	err = json.Unmarshal(bytes, todo)
	if err != nil {
		todo.Log.Errorf(msgErrorBrokenFile, path)
		return fmt.Errorf(msgErrorBrokenFile, path)
	}
	return nil
}
