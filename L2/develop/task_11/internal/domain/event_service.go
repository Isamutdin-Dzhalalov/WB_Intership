// /internal/domain/event_service.go

package domain

import (
    "errors"
    "time"
)

// Глобальная переменная для хранения всех событий.
var events = []Event{}
var currentID = 1
// Ф-ция для создания нового события.
func CreateEvent(userID int, title string, date time.Time, duration time.Duration, notes string) (Event, error) {
	// Проверяем корректность входных данных.
    if userID == 0 || title == "" || date.IsZero() {
        return Event{}, errors.New("invalid input")
    }
    event := Event{
        ID:       currentID,
        UserID:   userID,
        Title:    title,
        Date:     date,
        Duration: duration,
        Notes:    notes,
    }
    currentID++ // Увеличиваем текущий ID для следующего события.
    events = append(events, event)
    return event, nil
}

// Ф-ция для обновления события.
func UpdateEvent(id, userID int, title string, date time.Time, duration time.Duration, notes string) (Event, error) {
	// Итерируемся по всем событиям.
    for i, event := range events {
		// Если находим событие с нужным ID.
        if event.ID == id {
		// Обновляем событие.
            events[i] = Event{
                ID: id,
                UserID: userID,
                Title: title,
                Date: date,
                Duration: duration,
                Notes: notes,
            }
            return events[i], nil
        }
    }
	// Если событие с указанным ID не найдено, возвращаем ошибку.
    return Event{}, errors.New("event not found")
}

// Ф-ция для удаления события по ID.
func DeleteEvent(id int) error {
	// Итерируемся по всем событиям.
    for i, event := range events {
		// Если находим событие с нужным ID.
        if event.ID == id {
			// Удаляем это событие.
            events = append(events[:i], events[i+1:]...)
            return nil
        }
    }
	// Если событие с указанным ID не найдено, возвращаем ошибку.
    return errors.New("event not found")
}

/* Ф-ция для получения событий за указанный день.
	Итерируемся по событиям, если событие произошло в указанный день,
	добавляем ситуацию в result и возвращаем. */
func GetEventsForDay(date time.Time) ([]Event, error) {
    var result []Event
    for _, event := range events {
        if event.Date.Year() == date.Year() && event.Date.YearDay() == date.YearDay() {
            result = append(result, event)
        }
    }
    return result, nil
}
// Аналогично делаем и с неделей и с месяцем.
func GetEventsForWeek(date time.Time) ([]Event, error) {
    var result []Event
	// Определяем начало и конец недели.
    startOfWeek := date.AddDate(0, 0, -int(date.Weekday()))
    endOfWeek := startOfWeek.AddDate(0, 0, 7)

    for _, event := range events {
        if event.Date.After(startOfWeek) && event.Date.Before(endOfWeek) {
            result = append(result, event)
        }
    }
    return result, nil
}

func GetEventsForMonth(date time.Time) ([]Event, error) {
    var result []Event
	// Определяем начало и конец месяца.
    startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
    endOfMonth := startOfMonth.AddDate(0, 1, 0)

    for _, event := range events {
        if event.Date.After(startOfMonth) && event.Date.Before(endOfMonth) {
            result = append(result, event)
        }
    }
    return result, nil
}

