package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"user-management-service/pkg/handler"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

var Tmpl *template.Template
var DB *sql.DB
var Store = sessions.NewCookieStore([]byte("ums"))

const (
	Host     string = "localhost"
	Port     int    = 5432
	User     string = "postgres"
	Password string = "nourian1999"
	Dbname   string = "ums"
)

func init() {
	Tmpl, _ = template.ParseGlob("template/*.html")
	//setup sessions
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 3,
		HttpOnly: true,
	}
}

func initDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func main() {
	initDB()


	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Tmpl.ExecuteTemplate(w, "home.html", nil)
	})
	gRouter.HandleFunc("/register", handler.RegisterPage(DB, Tmpl)).Methods("GET")
	gRouter.HandleFunc("/register", handler.RegisterHandler(DB, Tmpl)).Methods("POST")
	http.ListenAndServe(":8080", gRouter)

}
