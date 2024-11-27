package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	gRouter := mux.NewRouter()
	gRouter.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":3000", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NimiraTech | Home Page")
}