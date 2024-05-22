// /internal/middleware/logging.go

package middleware

import (
    "log"
    "net/http"
    "time"
)

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    // Записываем текущее время перед обработкой запроса.
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
		// Передаем запрос следующему обработчику в цепочке.
        next.ServeHTTP(w, r)
        log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
    })
}

