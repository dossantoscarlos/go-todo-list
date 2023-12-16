package main

import (
	"fmt"
	"log"
	"net/http"

	"go-todo-list/src/api/routes"
)

func main() {
	fmt.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", routes.Router()))
}
