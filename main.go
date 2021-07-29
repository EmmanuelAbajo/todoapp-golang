package main

import (
	"log"
	"net/http"
	"os"
	ct "todo-app/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	d "todo-app/domain"
)


func main()  {
	conn := d.Connection()
	defer conn.Close()
	
	_ = godotenv.Load("conf.env")
	
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/todo", ct.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo", ct.GetAllTodosHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", ct.GetTodoHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", ct.DeleteTodoHandler).Methods("DELETE")

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(port,router); err != nil {
		log.Fatal(err)
	}
}