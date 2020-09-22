package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Marlos-Rodriguez/Go-PostgreSQL/src/helpers"
	"github.com/Marlos-Rodriguez/Go-PostgreSQL/src/models"
	"github.com/gorilla/mux"
)

//Data ...
type Data struct {
	Success bool          `json: "success"`
	Data    []models.Todo `json: "data"`
	Errors  []string      `json: "errors"`
}

//CreateTodo ...
func CreateTodo(w http.ResponseWriter, req *http.Request) {
	bodyTodo, success := helpers.DecodeBody(req)

	if success != true {
		http.Error(w, "Could not decode Body", http.StatusBadRequest)
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)

	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	todo, success := models.Insert(bodyTodo.Description)

	if success != true {
		data.Errors = append(data.Errors, "could not create todo")
	}

	data.Success = true

	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

//GetTodo ...
func GetTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var todo models.Todo
	var success bool

	todo, success = models.Get(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "todo no found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//GetTodos ...
func GetTodos(w http.ResponseWriter, req *http.Request) {
	var todos []models.Todo = models.GetAll()

	var data = Data{true, todos, make([]string, 0)}

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//DeleteTodo ...
func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data = Data{Errors: make([]string, 0)}

	todo, success := models.Delete(id)

	if success != true {
		data.Errors = append(data.Errors, "could not delete todo")
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//UpdateTodo ...
func UpdateTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	bodyTodo, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "Could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	todo, success := models.Update(id, bodyTodo.Description)
	if success != true {
		data.Errors = append(data.Errors, "could not update todo")
	}

	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}
