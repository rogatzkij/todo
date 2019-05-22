package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
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
	ok := ToDoDatabase.WriteUser(inputLogin, inputPass, inputEmail)
	if ok {
		// при удачной регистрации сразу же входим
		ToDoDatabase.Log.Infof("Successful registration: %s %s", inputLogin, inputEmail)
		http.SetCookie(w, cookieGen(inputLogin))
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	} else {
		// TODO неудачная регистация
	}
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
		ToDoDatabase.Log.Infof("Successful login: %s %s", user.Login, user.Email)
		http.SetCookie(w, cookieGen(user.Login))
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	} else {
		ToDoDatabase.Log.Errorf("Unsuccessful login: %s %s", user.Login, user.Email)
		// TODO неудачный вход
	}
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
	switch r.Method {
	case "GET":
		dashboardPageGET(w, r)
	case "POST":

	case "PUT":
		dashboardPagePUT(w, r)
	case "DELETE":
		dashboardPageDELETE(w, r)
	default:
		return
	}
}

func dashboardPageGET(w http.ResponseWriter, r *http.Request) {
	// узнаем имя из кук
	cookie, _ := r.Cookie("session_id")
	username, _ := ToDoDatabase.SearchUserByCookie(cookie.Value)

	// читаем дела
	tasks, _ := ToDoDatabase.GetAllTasks(username)

	// вставляем в шаблон
	tmpl := template.Must(template.ParseFiles(sTEMPLATE_FLOADER + "/dashboard.html"))
	tmpl.Execute(w, tasks)

	w.WriteHeader(http.StatusOK)
}

//добавляем новое дело или помечаем дело сделанным
func dashboardPagePUT(w http.ResponseWriter, r *http.Request) {

	// узнаем имя из кук
	cookie, _ := r.Cookie("session_id")
	login, _ := ToDoDatabase.SearchUserByCookie(cookie.Value)

	if taskID := r.FormValue("id"); taskID != "" {
		//помечаем дело сделанным
		ok := ToDoDatabase.taskDone(taskID, login)
		if ok == false {
			ToDoDatabase.Log.Errorf("Can't done task #%s by user:%s", taskID, login)
		} else {
			ToDoDatabase.Log.Infof("Done task #%s by user:%s", taskID, login)
		}
	} else {
		//добавляем новое дело
		task := Task{}

		task.Login = login
		task.Title.Scan(r.FormValue("title"))
		task.Description.Scan(r.FormValue("description"))
		task.DateEnd.Scan(r.FormValue("date"))

		ok := ToDoDatabase.AddTask(task)
		if ok == true {
			w.WriteHeader(http.StatusOK)
			ToDoDatabase.Log.Infof("Adding task successful, login: %s", login)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			ToDoDatabase.Log.Errorf("Adding task unsuccessful, login: %s", login)
		}
	}

}

// страница с делами - удаление дел
func dashboardPageDELETE(w http.ResponseWriter, r *http.Request) {

	// узнаем имя из кук
	cookie, _ := r.Cookie("session_id")
	login, _ := ToDoDatabase.SearchUserByCookie(cookie.Value)

	taskID := r.FormValue("id")

	ok := ToDoDatabase.DeleteTask(login, taskID)
	if ok == true {
		w.WriteHeader(http.StatusOK)
		ToDoDatabase.Log.Infof("Deleteing task successful, login: %s", login)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		ToDoDatabase.Log.Errorf("Deleteing task unsuccessful, login: %s", login)
	}
}

// архив дел
func archivePageGET(w http.ResponseWriter, r *http.Request) {
	// узнаем имя из кук
	cookie, _ := r.Cookie("session_id")
	username, _ := ToDoDatabase.SearchUserByCookie(cookie.Value)

	// читаем дела
	tasks, _ := ToDoDatabase.GetArchiveTasks(username)

	// вставляем в шаблон
	tmpl := template.Must(template.ParseFiles(sTEMPLATE_FLOADER + "/archive.html"))
	tmpl.Execute(w, tasks)

	w.WriteHeader(http.StatusOK)
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
			case "/archive":
				fallthrough
			case "/dashboard":
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

		} else { // если ок
			ToDoDatabase.Log.Infof("auth accepted: %s", r.URL.Path)
			switch r.URL.Path {
			case "/archive":
				//return
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
	userMux.HandleFunc("/archive", archivePageGET)
	userMux.HandleFunc("/registration", registrationPage)
	userMux.HandleFunc("/", mainPage)

	fullMux := http.NewServeMux()
	fullMux.Handle("/", authMiddleware(userMux))
	fullMux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	ToDoDatabase.Log.Info("starting server at :80")
	log.Fatal(http.ListenAndServe(":8081", fullMux))
}
