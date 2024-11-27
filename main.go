package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:password1234@(127.0.0.1:4444)/htmx-go?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}	
}

func main() {
	initDB()
	defer db.Close()
	gRouter := mux.NewRouter()
	gRouter.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":3000", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NimiraTech | Home Page")
}