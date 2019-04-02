package main

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

func (todo *DBtodo) writeUser(login, pswd, email string) error {

	hashBytes := (md5.Sum([]byte(pswd)))
	hash := fmt.Sprintf("%x", hashBytes)

	debug(fmt.Sprintf("Хэш %s %s", pswd, hash))

	_, err := todo.database.Exec(sWRITEUSER, login, email, hash)
	if err != nil {
		return fmt.Errorf("Неудачная запись пользователя %s", login)
	}

	return nil
}

/*
	Получить пользователя
*/

const (
	sGETUSER = `SELECT * FROM E1_Users WHERE login=? OR email=?`
)

func (todo *DBtodo) getUser(login, pswd, email string) (User, bool) {
	var usr User

	err := todo.database.Get(&usr, sGETUSER, login, email)

	if err != nil {
		debug(fmt.Sprintf("Неправильный login(%s) ", login))
		return User{}, false
	}

	hashBytes := (md5.Sum([]byte(pswd)))
	hash := fmt.Sprintf("%x", hashBytes)

	if usr.Hash != hash {
		debug(fmt.Sprintf("Неправильный пароль(%s)", pswd))
		return User{}, false
	}

	return usr, true
}
