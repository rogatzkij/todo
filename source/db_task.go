package main

import (
	"database/sql"
)

type Task struct {
	IDTask      int            `db:"idTask"`      //INTEGER AUTO_INCREMENT,
	Login       string         `db:"login"`       // VARCHAR(20) NOT NULL,
	Title       sql.NullString `db:"title"`       //INTEGER NULL,
	Description sql.NullString `db:"description"` //VARCHAR(20) NULL,
	Defer       sql.NullBool   `db:"defer"`       // INTEGER NULL,
	Done        sql.NullBool   `db:"done"`        // INTEGER NULL,
	DateEnd     sql.NullString `db:"dateEnd"`     // DATE NULL,
}

type TasksToTemplate struct {
	Today    []Task
	Tomorrow []Task
	Soon     []Task
}

// добавить
func (todo *DBtodo) AddTask(task Task) bool {
	_, err := todo.Database.Exec(
		`INSERT INTO E3_Tasks (login, title, description, defer, done, dateEnd) VALUES (?, ?, ?, ?, ?, ?)`,
		task.Login,
		task.Title.String,
		task.Description.String,
		task.Defer.Bool,
		task.Done.Bool,
		task.DateEnd.String,
	)
	if err != nil {
		todo.Log.Errorf("Error add task")
		return false
	}
	return true
}

// удалить
func (todo *DBtodo) DeleteTask(login string, IDTask string) bool {
	_, err := todo.Database.Exec(
		`DELETE FROM E3_Tasks WHERE login = ? AND IDtask = ?`,
		login,
		IDTask,
	)
	if err != nil {
		todo.Log.Errorf("Error delete task: user:%s ID:%s ",
			login,
			IDTask,
		)
		return false
	}

	return true
}

// получить
func (todo *DBtodo) GetAllTasks(login string) (TasksToTemplate, bool) {

	today, ok := todo.GetTodayTasks(login)
	if ok == false {
		todo.Log.Error("Can't get today tasks, login:%s", login)
		return TasksToTemplate{}, false
	}

	tomorrow, ok := todo.GetTomorrowTasks(login)
	if ok == false {
		todo.Log.Error("Can't get tomorrow tasks, login:%s", login)
		return TasksToTemplate{}, false
	}

	soon, ok := todo.GetSoonTasks(login)
	if ok == false {
		todo.Log.Error("Can't get tomorrow tasks, login:%s", login)
		return TasksToTemplate{}, false
	}

	return TasksToTemplate{
		Today:    today,
		Tomorrow: tomorrow,
		Soon:     soon,
	}, true
}

// получить сегодня
func (todo *DBtodo) GetTodayTasks(login string) ([]Task, bool) {
	var task []Task

	err := todo.Database.Select(&task,
		"SELECT * FROM `E3_Tasks` WHERE DATEDIFF(dateEnd, NOW())=0 AND done=0 AND login = ?",
		login,
	)

	if err != nil {
		//	debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return []Task{}, false
	}

	return task, true
}

// получить завтра
func (todo *DBtodo) GetTomorrowTasks(login string) ([]Task, bool) {
	var task []Task

	err := todo.Database.Select(&task,
		"SELECT * FROM `E3_Tasks` WHERE DATEDIFF(dateEnd, NOW())=1 AND done=0 AND login = ?",
		login,
	)

	if err != nil {
		//	debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return []Task{}, false
	}

	return task, true
}

// получить скоро
func (todo *DBtodo) GetSoonTasks(login string) ([]Task, bool) {
	var task []Task

	err := todo.Database.Select(&task,
		"SELECT * FROM `E3_Tasks` WHERE DATEDIFF(dateEnd, NOW())>1 AND done=0 AND login = ?",
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

func (todo *DBtodo) taskDone(IDTask string, login string) bool {
	_, err := todo.Database.Exec(
		"UPDATE E3_Tasks SET done=1 WHERE idTask=? AND login=?",
		IDTask,
		login,
	)

	if err != nil {
		return false
	}
	return true
}
