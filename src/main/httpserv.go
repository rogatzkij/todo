package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	sTEMPLATE_FLOADER = "./templates" // путь с шаблонами страниц
)

// ============================
// ||       Хэндлеры         ||
// ============================

// страница регистрации
func registrationPage(w http.ResponseWriter, r *http.Request) {

	path := sTEMPLATE_FLOADER + "/registration.html"

	registrationForm, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Файл %s не найден\n", path)
		return
	}

	if r.Method != http.MethodPost {
		w.Write(registrationForm)
		return
	}

	inputLogin := r.FormValue("login")
	inputPass := r.FormValue("password")
	inputEmail := r.FormValue("email")

	fmt.Fprintln(w, "you enter: ", inputLogin, inputPass, inputEmail)
	ToDoDatabase.WriteUser(inputLogin, inputPass, inputEmail)
}

// страница авторизации
func loginPage(w http.ResponseWriter, r *http.Request) {

	path := sTEMPLATE_FLOADER + "/login.html"

	loginForm, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Файл %s не найден\n", path)
		return
	}

	if r.Method != http.MethodPost {
		w.Write(loginForm)
		return
	}

	inputLogin := r.FormValue("login")
	inputPass := r.FormValue("password")

	user, ok := ToDoDatabase.GetUser(inputLogin, inputPass, inputLogin)
	if ok {
		ToDoDatabase.Log.Infof("you are: %s %s", user.Login, user.Email)
		http.SetCookie(w, cookieGen(user.Login))
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}
	fmt.Fprintln(w, "Все плохо")

}

// генерация куки
func cookieGen(user string) *http.Cookie {

	expiration := time.Now().Add(time.Hour) // время жизни куки

	hashBytes := (md5.Sum([]byte(user + time.Now().String())))
	hash := fmt.Sprintf("%x", hashBytes)

	//запись куки в бд
	ToDoDatabase.WriteCookie(user, hash)

	cookie := http.Cookie{Name: "session_id", Value: hash, Expires: expiration}
	return &cookie
}

// главная страница
func mainPage(w http.ResponseWriter, r *http.Request) {
	path := sTEMPLATE_FLOADER + "/main.html"

	loginForm, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Файл %s не найден\n", path)
		return
	}

	w.Write(loginForm)
	return

}

// страница с делами
func dashboardPage(w http.ResponseWriter, r *http.Request) {
	path := sTEMPLATE_FLOADER + "/dashboard.html"

	loginForm, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Файл %s не найден\n", path)
		return
	}

	w.Write(loginForm)
	return
}

// ============================
// ||       Мидлвары         ||
// ============================

// мидлвар авторизации
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ToDoDatabase.Log.Info("authMiddleware", r.URL.Path)

		ok := false
		cookie, err := r.Cookie("session_id")
		if err == nil {
			_, ok = ToDoDatabase.SearchUserByCookie(cookie.Value)
		}

		if ok == false { // если не ок
			ToDoDatabase.Log.Errorf("auth failed: %s", r.URL.Path)
			switch r.URL.Path {
			case "/dashboard":
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

		} else { // если ок
			ToDoDatabase.Log.Infof("auth accepted: %s", r.URL.Path)
			switch r.URL.Path {
			case "/dashboard":
				//return
			case "/":
				fallthrough
			case "/login":
				fallthrough
			case "/registration":
				http.Redirect(w, r, "/dashboard", http.StatusFound)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// ============================
// ||         Main           ||
// ============================

// слушаем порт
func serv() {
	userMux := http.NewServeMux()
	userMux.HandleFunc("/login", loginPage)
	userMux.HandleFunc("/dashboard", dashboardPage)
	userMux.HandleFunc("/registration", registrationPage)
	userMux.HandleFunc("/", mainPage)

	fullMux := http.NewServeMux()
	fullMux.Handle("/", authMiddleware(userMux))
	fullMux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	ToDoDatabase.Log.Info("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", fullMux))
}
