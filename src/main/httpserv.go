package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	sTEMPLATE_FLOADER = "./templates"
)

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

func serv() {
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/", registrationPage)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
