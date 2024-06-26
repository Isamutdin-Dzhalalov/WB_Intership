// /pkg/json_helpers.go

package pkg

import (
    "encoding/json"
    "net/http"
)

func WriteJSON(w http.ResponseWriter, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(v)
}

