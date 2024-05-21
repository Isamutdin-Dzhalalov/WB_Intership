package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Структуры для объектов доменной области
type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

// Вспомогательные функции для сериализации и десериализации
func serializeEvent(event Event) ([]byte, error) {
	return json.Marshal(event)
}

func parseEventParams(r *http.Request) (*Event, error) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	return &event, err
}

// HTTP обработчики для методов API
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := parseEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Логика создания события
	result := map[string]string{"result": "event created"}
	json.NewEncoder(w).Encode(result)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := parseEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Логика обновления события
	result := map[string]string{"result": "event updated"}
	json.NewEncoder(w).Encode(result)
}

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Регистрация обработчиков маршрутов
	http.Handle("/create_event", loggingMiddleware(http.HandlerFunc(createEventHandler)))
	http.Handle("/update_event", loggingMiddleware(http.HandlerFunc(updateEventHandler)))

	// Запуск HTTP-сервера на порту :8080
	log.Println("HTTP-сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

