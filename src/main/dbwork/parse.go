package dbwork

import (
	"encoding/json"
	"io/ioutil"
)

// чтение настроек из конфиг файла
func (todo *DBtodo) ParseSQLSettings(path string) {

	bytes, err := ioutil.ReadFile(path)
	if err != nil { // если файла нет создаем с настройками поумолчанию
		todo.Log.Error("Не удалось открыть файл " + path)

		todo.DriverName = "mysql"
		todo.Login = "root"
		todo.Password = ""
		todo.Adress = "localhost"
		todo.Port = 3306
		todo.DatabaseName = "todo"

		bytes, _ := json.Marshal(todo)

		err = ioutil.WriteFile(path, bytes, 0644)
		if err != nil {
			todo.Log.Error("Не удалось создать файл " + path)
		}

		todo.Log.Error("Создан файл с начальными настройками " + path)
		return
	}

	err = json.Unmarshal(bytes, todo)
	if err != nil {
		todo.Log.Error("Файл испорчен " + path)
	}

}

func (t *DBtodo) ParseTemplateSettings(path string) {

}
