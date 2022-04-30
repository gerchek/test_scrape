package main

import (
	"fmt"
	"html/template"
	"net/http"
	"scraper_trendyol/handler"
	"scraper_trendyol/pkg/logging"
	"scraper_trendyol/setup"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

func init() {

	logrus.Infoln("init function")

	// init .env
	err := setup.LoadEnv()
	if err != nil {
		logging.Error("error, can not load .env file: ", err.Error())
		panic(err.Error())
	}
	// setup logrus time setup
	setup.LogrusSetup()

	// logging setup
	logging.Setup()

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/init-scraper", handler.InitScraper)

	r.HandleFunc("/parse-link", handler.ParseLink)

	r.HandleFunc("/parse-excel", handler.ParseExcel)

	r.HandleFunc("/init-updater", handler.InitUpdater)

	fs := http.FileServer(http.Dir("static"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/index", Dashboard).Methods("GET")

	r.HandleFunc("/login", Login).Methods("GET")

	r.HandleFunc("/account/login", Account_Login)
	r.HandleFunc("/account/logout", Account_Logout)

	// -----------------------------------------------------------------------------crud
	r.HandleFunc("/GetExcelData", handler.GetExcel)

	logging.Info("[info] start http server listening 9000")
	err := http.ListenAndServe(":9000", r)

	if err != nil {
		logging.Error("error: ", err)
	}

}

var store = sessions.NewCookieStore([]byte("mysession"))

func Dashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	username := session.Values["username"]
	password := session.Values["password"]
	fmt.Println("username ", username)
	fmt.Println("password ", password)
	//
	if username == "abs" && password == "123" {
		data := map[string]interface{}{
			"username": username,
		}
		tmpl := template.Must(template.ParseFiles("./views/pages/index.html", "./views/layouts/default.html", "./views/partials/vertical_menu.html"))
		tmpl.ExecuteTemplate(w, "default", data)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./views/pages/login.html", "./views/layouts/default.html")
	tmpl.ExecuteTemplate(w, "default", "users")
}

func Account_Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username == "abs" && password == "123" {
		session, _ := store.Get(r, "mysession")
		session.Values["username"] = username
		session.Values["password"] = password
		// fmt.Println("username ", username)
		session.Save(r, w)
		http.Redirect(w, r, "/index", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func Account_Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	session.Values["username"] = nil
	username := session.Values["username"]
	fmt.Println("logout username ", username)

	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
