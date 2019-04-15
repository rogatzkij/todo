package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	sTEMPLATE_FLOADER = "./templates" // путь с шаблонами страниц
)

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
		fmt.Fprintln(w, "Все хорошо, ты: ", user.Login, user.Email, user.Hash)
		return
	}
	fmt.Fprintln(w, "Все плохо")

}

// главная страница
func mainPage(w http.ResponseWriter, r *http.Request) {

}

// мидлвар авторизации
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ToDoDatabase.Log.Info("adminAuthMiddleware", r.URL.Path)

		_, err := r.Cookie("session_id")
		// учебный пример! это не проверка авторизации!
		if err != nil {
			fmt.Println("no auth at", r.URL.Path)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// слушаем порт
func serv() {
	r := mux.NewRouter()

	r.HandleFunc("/login", loginPage)
	r.HandleFunc("/", registrationPage)

	ToDoDatabase.Log.Info("starting server at :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
