package main

import (
	"fmt"
	"net/http"
)
//definição de tipos para as varieaveis armazenadas 
type Task struct{
	ID 	int 	`json:"id"`
	Title	string `json:"title"`
	Done bool `json:"done"`
}
var task []Task;

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


func createTask(w http.ResponseWriter, r *http.Request){
	if 
}
