package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/dossantoscarlos/go-todo-list/api/models"
	"github.com/dossantoscarlos/go-todo-list/api/service"

	"github.com/gorilla/mux"
)

func IndexTaskHandler(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{"Content-Type": "application/json"}
	tasks, _ := service.AllTask()
	data := map[string]interface{}{"data": tasks}
	responseJson(w, http.StatusOK, data, headers)
}

func ShowTodoListHandler(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{"Content-Type": "application/json"}
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Default().Output(log.LUTC, err.Error())
		http.Error(w, messageErrorHandler("ID invalido!"), http.StatusBadRequest)
		return
	}

	task, err := service.FindOneTask(id)
	if err != nil {
		log.Default().Output(log.LUTC, err.Error())
		http.Error(w, messageErrorHandler("Task não existe"), http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"data": map[string]interface{}{
			"task_id":     task.ID,
			"task_name":   task.Name,
			"task_status": task.Status,
		},
	}

	responseJson(w, http.StatusOK, data, headers)
}

func SaveTodoListHandler(response http.ResponseWriter, request *http.Request) {
	var task models.Task

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Default().Output(log.LUTC, err.Error())
		http.Error(response, messageErrorHandler("Erro ao ler requisição"), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Default().Output(log.LUTC, err.Error())
		http.Error(response, messageErrorHandler("Dados mal formatado"), http.StatusBadRequest)
		return
	}

	taskID, err := service.SaveTask(task)
	if err != nil {
		log.Default().Output(log.LUTC, err.Error())
		http.Error(response, messageErrorHandler(err.Error()), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{"message": "Task created"}
	scheme := request.Header.Get("X-Forwarded-Proto")

	if scheme == "" {
		scheme = "http"
	}

	recurso := fmt.Sprintf("%s://%s/api/tasks/%d", scheme, request.Host, taskID)
	headers := map[string]string{"Location": recurso, "Content-Type": "application/json"}

	responseJson(response, http.StatusCreated, data, headers)
}

func UpdateTodoListHandler(response http.ResponseWriter, request *http.Request) {
	var todoList models.Task
	param := mux.Vars(request)
	id := param["id"]
	data := map[string]interface{}{"id": id}
	headers := map[string]string{"Content-Type": "application/json"}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Default().Output(log.LUTC, err.Error())
		http.Error(response, messageErrorHandler("Erro ao ler o corpo da requisição"), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &todoList)
	if err != nil {
		http.Error(response, messageErrorHandler("Dados mal formatado!"), http.StatusBadRequest)
		return
	}
	responseJson(response, http.StatusAccepted, data, headers)
}

func DeleteTodoListHandler(response http.ResponseWriter, request *http.Request) {
	data := map[string]interface{}{"message": "Ação executada com sucesso"}
	headers := map[string]string{"Content-Type": "application/json"}
	responseJson(response, http.StatusAccepted, data, headers)
}

func responseJson(response http.ResponseWriter, status int, data interface{}, headers map[string]string) {

	for key, value := range headers {
		response.Header().Set(key, value)
	}

	response.WriteHeader(status)

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Erro ao serializar para JSON:", err)
		return
	}

	response.Write(jsonData)
}

func messageErrorHandler(info string) string {
	message := map[string]string{"message": info}
	messageJson, _ := json.Marshal(message)
	return string(messageJson)
}
