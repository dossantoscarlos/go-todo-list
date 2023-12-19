package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dossantoscarlos/go-todo-list/api/routes"
)

func main() {
	fmt.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", routes.Router()))
}
