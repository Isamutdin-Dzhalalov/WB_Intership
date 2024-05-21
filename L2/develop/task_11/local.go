package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Обработчик для маршрута "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Привет, мир!"})
	})

	// Запуск HTTP-сервера на порту :8080
	log.Println("HTTP-сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

