package dbwork

import "fmt"

type CookieTodo struct {
	IDSession int    `db:"idSession"`
	Cookie    string `db:"cookie"`
	Login     string `db:"login"`
}

/*
	Проверить наличие cookie в БД
*/
const (
	sUSERBYCOOKIE = `SELECT * FROM E2_Session WHERE cookie=?`
)

func (todo *DBtodo) SearchUserByCookie(cookie string) (string, bool) {
	var ck CookieTodo

	err := todo.Database.Select(&ck, sUSERBYCOOKIE, cookie)

	if err != nil {
		//		Debug(fmt.Sprintf("Нет куки %s", cookie))
		return "", false
	}

	//	Debug(fmt.Sprintf("Кука найдена %s - @%s", cookie, ck.Cookie))
	return ck.Cookie, true
}

/*
	Записать cookie
*/
const (
	sWRITECOOKIE = `INSERT INTO E2_Session (login, cookie) VALUES (?, ?)`
)

func (todo *DBtodo) WriteCookie(login, cookie string) error {

	_, err := todo.Database.Exec(sWRITECOOKIE, login, cookie)
	if err != nil {
		return fmt.Errorf("Неудачная запись %s", cookie)
	}

	return nil
}

/*
	Удалить cookie из БД
*/
const (
	sDELETECOOKIE = `DELETE FROM E2_Session where cookie=?`
)

func (todo *DBtodo) DeleteCookie(cookie string) error {

	_, err := todo.Database.Exec(sDELETECOOKIE, cookie)
	if err != nil {
		return fmt.Errorf("Неудачное удаление %s", cookie)
	}

	return nil
}

/*
	Удалить все cookie для пользователя
*/
const (
	sDELETEALLCOOKIE = `DELETE FROM E2_Session where login=?`
)

func (todo *DBtodo) DeleteAllCookie(login string) error {

	_, err := todo.Database.Exec(sDELETEALLCOOKIE, login)
	if err != nil {
		return fmt.Errorf("Неудачное удаление cookie для %s", login)
	}

	return nil
}
