// /cmd/main.go

package main

import (
    "log"
    "net/http"
    "task_11/internal/handlers"
    "task_11/internal/middleware"
)

func main() {
	// Создаём маршрутизатор.
    mux := http.NewServeMux()
	// Обработчики для разных маршрутов, обрабатывают http-запросы.
    mux.HandleFunc("/create_event", handlers.CreateEvent)
    mux.HandleFunc("/update_event", handlers.UpdateEvent)
    mux.HandleFunc("/delete_event", handlers.DeleteEvent)
    mux.HandleFunc("/events_for_day", handlers.EventsForDay)
    mux.HandleFunc("/events_for_week", handlers.EventsForWeek)
    mux.HandleFunc("/events_for_month", handlers.EventsForMonth)

	// Для логирования.
    loggedMux := middleware.Logging(mux)

	// Запускаем сервер.
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", loggedMux)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

