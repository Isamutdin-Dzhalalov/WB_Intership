package handlers

import (
  "encoding/json"
    "net/http"
    "task_11/internal/domain"
    "task_11/pkg"
    "time"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Парсинг данных из тела Post-запроса.
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

	// Проверяем валидность данных.
    userID, err := pkg.ParseInt(r.PostForm, "user_id")
    if err != nil {
        http.Error(w, "Invalid user_id", http.StatusBadRequest)
        return
    }

	// Извлекаем значение из данных формы. 
    title := r.PostForm.Get("title")
    if title == "" {
        http.Error(w, "Missing title", http.StatusBadRequest)
        return
    }

    date, err := pkg.ParseDate(r.PostForm, "date")
    if err != nil {
        http.Error(w, "Invalid date", http.StatusBadRequest)
        return
    }

	// Извлекаем значение из данных формы. 
    durationStr := r.PostForm.Get("duration")
    if durationStr == "" {
        http.Error(w, "Missing duration", http.StatusBadRequest)
        return
    }

	// Парсим строку.
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        http.Error(w, "Invalid duration", http.StatusBadRequest)
        return
    }

    notes := r.PostForm.Get("notes")

    event, err := domain.CreateEvent(userID, title, date, duration, notes)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

	// Устанавливает заголовок Content-Type на application/json.
    w.Header().Set("Content-Type", "application/json")
	// Кодирует результат в JSON и отправляет клиенту.
    json.NewEncoder(w).Encode(map[string]interface{}{"result": event})
}

// Аналогично ф-ции Create.
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    eventID, err := pkg.ParseInt(r.PostForm, "id")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    userID, err := pkg.ParseInt(r.PostForm, "user_id")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    title := r.PostFormValue("title")
    date, err := pkg.ParseDate(r.PostForm, "date")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    durationStr := r.PostFormValue("duration")
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    notes := r.PostFormValue("notes")
    event, err := domain.UpdateEvent(eventID, userID, title, date, duration, notes)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

    pkg.WriteJSON(w, map[string]interface{}{"result": event})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    eventID, err := pkg.ParseInt(r.PostForm, "id")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = domain.DeleteEvent(eventID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

    pkg.WriteJSON(w, map[string]interface{}{"result": "event deleted"})
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	// Проверяет, что метод запроса GET.
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	// Извлекает дату из строки запроса.
    date, err := pkg.ParseDate(r.URL.Query(), "date")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	
	// Ф-ция для получения событий за указанный день.
    events, err := domain.GetEventsForDay(date)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

    pkg.WriteJSON(w, map[string]interface{}{"result": events})
}

// Аналогично, как и ф-ция EventsForDay().
func EventsForWeek(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    date, err := pkg.ParseDate(r.URL.Query(), "date")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    events, err := domain.GetEventsForWeek(date)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

    pkg.WriteJSON(w, map[string]interface{}{"result": events})
}
func EventsForMonth(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    date, err := pkg.ParseDate(r.URL.Query(), "date")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    events, err := domain.GetEventsForMonth(date)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

    pkg.WriteJSON(w, map[string]interface{}{"result": events})
}

