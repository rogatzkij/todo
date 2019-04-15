package dbwork

import (
	"crypto/md5"
	"fmt"
)

type User struct {
	Login string `db:"login"`
	Email string `db:"email"`
	Hash  string `db:"hash"`
}

/*
	Создание учетки
*/
const (
	sWRITEUSER = `INSERT INTO E1_Users (login, email, hash) VALUES (?, ?, ?)`
)

func (todo *DBtodo) WriteUser(login, pswd, email string) (bool, error) {

	hashBytes := (md5.Sum([]byte(pswd)))
	hash := fmt.Sprintf("%x", hashBytes)

	//debug(fmt.Sprintf("Хэш %s %s", pswd, hash))

	// проверка на существование пользователя
	if exist, _ := todo.isExistUser(login, email); exist {
		return false, nil
	}

	_, err := todo.Database.Exec(sWRITEUSER, login, email, hash)
	if err != nil {
		return false, fmt.Errorf("Неудачная запись пользователя %s", login)
	}

	return true, nil
}

/*
	Получить пользователя
*/

const (
	sGETUSER = `SELECT * FROM E1_Users WHERE login=? OR email=?`
)

func (todo *DBtodo) GetUser(login, pswd, email string) (User, bool) {
	var usr User

	err := todo.Database.Get(&usr, sGETUSER, login, email)

	if err != nil {
		//	debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return User{}, false
	}

	hashBytes := (md5.Sum([]byte(pswd)))
	hash := fmt.Sprintf("%x", hashBytes)

	if usr.Hash != hash {
		//	debug(fmt.Sprintf("Неправильный пароль(%s)", pswd))
		return User{}, false
	}

	return usr, true
}

/*
	Получить пользователя
*/

const (
	sGETUSERCOUNT = `SELECT count(*) FROM E1_Users WHERE login=? OR email=?`
)

func (todo *DBtodo) isExistUser(login, email string) (bool, error) {
	var count int

	err := todo.Database.Get(&count, sGETUSERCOUNT, login, email)

	if err != nil {
		return false, fmt.Errorf("Ошибочка %s", login)
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
