package main

type ArchiveTasksToTemplate struct {
	Done    []Task
	NotDone []Task
}

func (todo *DBtodo) GetArchiveTasks(login string) (ArchiveTasksToTemplate, bool) {

	done, ok := todo.GetDoneTasks(login)
	if ok == false {
		todo.Log.Error("Can't get done tasks, login:%s", login)
		return ArchiveTasksToTemplate{}, false
	}

	notDone, ok := todo.GetNotDoneTasks(login)
	if ok == false {
		todo.Log.Error("Can't get not done tasks, login:%s", login)
		return ArchiveTasksToTemplate{}, false
	}

	return ArchiveTasksToTemplate{
		Done:    done,
		NotDone: notDone,
	}, true
}

func (todo *DBtodo) GetDoneTasks(login string) ([]Task, bool) {
	var task []Task

	err := todo.Database.Select(&task,
		"SELECT * FROM `E3_Tasks` WHERE done=1 AND login = ?",
		login,
	)

	if err != nil {
		//	debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return []Task{}, false
	}

	return task, true
}

func (todo *DBtodo) GetNotDoneTasks(login string) ([]Task, bool) {
	var task []Task

	err := todo.Database.Select(&task,
		"SELECT * FROM `E3_Tasks` WHERE DATEDIFF(dateEnd, NOW())<0 AND done=0 AND login = ?",
		login,
	)

	if err != nil {
		//	debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return []Task{}, false
	}

	return task, true
}
