package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type CreateEventRequest struct {
	Event Event `json:"event"`
}

type UpdateEventRequest struct {
	ID    string `json:"id"`
	Event Event `json:"event"`
}

type DeleteEventRequest struct {
	ID string `json:"id"`
}
func serializeEvent(event Event) ([]byte, error) {
	return json.Marshal(event)
}

func deserializeEvent(data []byte) (*Event, error) {
	var event Event
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseAndValidateCreateEventRequest(r *http.Request) (*CreateEventRequest, error) {
	var req CreateEventRequest
	err := json.NewDecoder(r.Body).Decode(&req.Event)
	if err!= nil {
		return nil, err
	}
	// Здесь можно добавить дополнительную валидацию, например, проверку даты
	return &req, nil
}

func parseAndValidateUpdateEventRequest(r *http.Request) (*UpdateEventRequest, error) {
	var req UpdateEventRequest
	err := json.NewDecoder(r.Body).Decode(&req.Event)
	if err!= nil {
		return nil, err
	}
	// Валидация здесь
	return &req, nil
}

func parseDeleteEventRequest(r *http.Request) (*DeleteEventRequest, error) {
	var req DeleteEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err!= nil {
		return nil, err
	}
	return &req, nil
}
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := parseAndValidateCreateEventRequest(r)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Логика создания события
	result := map[string]string{"result": "event created"}
	json.NewEncoder(w).Encode(result)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := parseAndValidateUpdateEventRequest(r)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Логика обновления события
	result := map[string]string{"result": "event updated"}
	json.NewEncoder(w).Encode(result)
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := parseDeleteEventRequest(r)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Логика удаления события
	result := map[string]string{"result": "event deleted"}
	json.NewEncoder(w).Encode(result)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Логика получения событий за день
	result := map[string]string{"result": "events for day"}
	json.NewEncoder(w).Encode(result)
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Логика получения событий за неделю
	result := map[string]string{"result": "events for week"}
	json.NewEncoder(w).Encode(result)
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Логика получения событий за месяц
	result := map[string]string{"result": "events for month"}
	json.NewEncoder(w).Encode(result)
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
func main() {
	http.Handle("/create_event", loggingMiddleware(http.HandlerFunc(createEventHandler)))
	http.Handle("/update_event", loggingMiddleware(http.HandlerFunc(updateEventHandler)))
	http.Handle("/delete_event", loggingMiddleware(http.HandlerFunc(deleteEventHandler)))
	http.Handle("/events_for_day", loggingMiddleware(http.HandlerFunc(eventsForDayHandler)))
	http.Handle("/events_for_week", loggingMiddleware(http.HandlerFunc(eventsForWeekHandler)))
	http.Handle("/events_for_month", loggingMiddleware(http.HandlerFunc(eventsForMonthHandler)))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Привет, мир!"})
	})
	http.ListenAndServe(":8080", nil)
}

