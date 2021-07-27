package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo-app/dto"
	"github.com/gorilla/mux"
)

var todos []dto.TodoResponseDTO

// CreateTodoHandler function
func CreateTodoHandler(res http.ResponseWriter, req *http.Request) {

	reqObj := &dto.TodoRequestDTO{}
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(reqObj); err != nil {
		http.Error(res, err.Error(),http.StatusBadRequest)
		return
	}
	

	// Set response
	resObj := &dto.TodoResponseDTO{}
	resObj.ID = len(todos) + 1;
	resObj.Name = reqObj.Name
	resObj.Content = reqObj.Content

	todos = append(todos, *resObj)
	
	// encode response
	res.Header().Set("Content-Type", "application/json ")
	if err := json.NewEncoder(res).Encode(resObj); err != nil {
		log.Printf("Can't encode %v - %s", res, err)
	}
}

// GetAllTodosHandler function
func GetAllTodosHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json ")
	json.NewEncoder(res).Encode(todos)
}

// GetTodoHandler function
func GetTodoHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json ")
	params := mux.Vars(req)
	for _, todo := range todos {
		if val, _ := strconv.Atoi(params["id"]); val == todo.ID {
			json.NewEncoder(res).Encode(todo)
			return
		}
	}
}

// DeleteTodoHandler function
func DeleteTodoHandler(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json ")
	params := mux.Vars(req)
	for index, todo := range todos {
		if val, _ := strconv.Atoi(params["id"]); val == todo.ID {
			todos = append(todos[:index], todos[index+1:]...)
		}
	}
}