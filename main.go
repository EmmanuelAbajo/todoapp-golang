package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	ct "todo-app/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"todo-app/domain"
)


func main()  {
	_ = godotenv.Load("conf.env")

	fmt.Println("Conecting to Database...")
	dbConn, err := sql.Open("mysql", domain.GetDbConnectionURL())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Conection to Database successful!")
	defer dbConn.Close()
	
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/todo", ct.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo", ct.GetAllTodosHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", ct.GetTodoHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", ct.DeleteTodoHandler).Methods("DELETE")

	fmt.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(port,router); err != nil {
		log.Fatal(err)
	}
}