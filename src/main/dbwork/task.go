package dbwork

import (
	"database/sql"
)

type Task struct {
	IDTask      int            `db:"idTask"`      //INTEGER AUTO_INCREMENT,
	Login       string         `db:"login"`       // VARCHAR(20) NOT NULL,
	Title       sql.NullString `db:"title"`       //INTEGER NULL,
	Description sql.NullString `db:"description"` //VARCHAR(20) NULL,
	Defer       sql.NullBool   `db:"defer"`       // INTEGER NULL,
	DateEnd     sql.NullString `db:"dateEnd"`     // DATE NULL,
}

// добавить
func (todo *DBtodo) AddTask(task Task) bool {
	_, err := todo.Database.Exec(
		`INSERT INTO E3_Tasks (login, title, description, defer, dateEnd) VALUES (?, ?, ?, ?, ?)`,
		task.Login,
		task.Title.String,
		task.Description.String,
		task.Defer.Bool,
		task.DateEnd.String,
	)
	if err != nil {
		todo.Log.Errorf("Error add task")
		return false
	}
	return true
}

/*
// удалить
func (todo *DBtodo) DeleteTask(login string, IDTask int) (Task, bool) {

}

// изменить
*/
// получить
func (todo *DBtodo) GetAllTasks(login string) ([]Task, bool) {
	var task []Task

	err := todo.Database.Select(&task,
		`SELECT * FROM E3_Tasks WHERE login=?`,
		login,
	)

	if err != nil {
		//	debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return []Task{}, false
	}

	return task, true
}

func (todo *DBtodo) GetTask(login string, IDTask int) (Task, bool) {
	var task Task

	err := todo.Database.Get(&task,
		`SELECT * FROM E3_Tasks WHERE login=? AND idTask=?`,
		login,
		IDTask,
	)

	if err != nil {
		//	debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return Task{}, false
	}

	return task, true
}
