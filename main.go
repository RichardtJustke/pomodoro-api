package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 🧠 Estrutura da Task
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// 🧠 "Banco fake" (memória)
var tasks []Task

func main() {
	// rota de teste
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	// rota de criação
	http.HandleFunc("/tasks", createTask)

	fmt.Println("Servidor rodando em http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

// 🚀 Criar Task (POST /tasks)
func createTask(w http.ResponseWriter, r *http.Request) {
	// 1. validar método
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. receber JSON
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	// 3. processar
	task.ID = len(tasks) + 1
	task.Done = false

	// 4. salvar
	tasks = append(tasks, task)

	// 5. responder
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
