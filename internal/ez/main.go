package ez

import (
	"encoding/json"
	"net/http"
	"service-status/config/ez"
)

func Routine(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(ez.Endpoints)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
