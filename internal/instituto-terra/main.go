package instituto_terra

import (
	"encoding/json"
	"net/http"
	"service-status/config/instituto-terra"
)

func Routine(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(instituto_terra.Endpoints)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
