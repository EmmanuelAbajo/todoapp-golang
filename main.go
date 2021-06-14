package main

import (
	"fmt"
	"log"
	"net/http"
	ct "todo-app/controller"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {

	db, err := sql.Open("mysql", "admin:password@tcp(127.0.0.1:6604)/testdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/todo", ct.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo", ct.GetAllTodosHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", ct.GetTodoHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", ct.DeleteTodoHandler).Methods("DELETE")

	fmt.Println("Server running on port 8090")
	if err := http.ListenAndServe(":8090",router); err != nil {
		log.Fatal(err)
	}
}