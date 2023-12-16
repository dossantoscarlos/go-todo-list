package routes

import (
	"go-todo-list/src/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/api/tasks", handlers.IndexTaskHandler).Methods("GET")
	route.HandleFunc("/api/tasks", handlers.SaveTodoListHandler).Methods("POST")
	route.HandleFunc("/api/tasks/{id}", handlers.ShowTodoListHandler).Methods("GET")
	route.HandleFunc("/api/tasks/{id}", handlers.UpdateTodoListHandler).Methods("PUT")
	route.HandleFunc("/api/tasks/{id}", handlers.DeleteTodoListHandler).Methods("DELETE")

	route.Use(corsMiddleware)

	return route
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(response, request)
	})
}
