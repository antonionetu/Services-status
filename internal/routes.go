package internal

import (
	"encoding/json"
	"net/http"
	"service-status/config/ez"
	"service-status/config/instituto-terra"
	"service-status/model"
)

type PathsStructure struct {
	ProjectUrl string
	Paths      []model.EndPoint
}

func GetAllPaths(w http.ResponseWriter, r *http.Request) {
	allPaths := []PathsStructure{
		{
			ProjectUrl: ez.Url,
			Paths:      ez.Endpoints,
		}, {
			ProjectUrl: instituto_terra.Url,
			Paths:      instituto_terra.Endpoints,
		},
	}

	jsonData, err := json.Marshal(allPaths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
