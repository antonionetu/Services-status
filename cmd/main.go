package main

import (
	"fmt"
	"net/http"
	"service-status/internal"

	"service-status/internal/ez"
	"service-status/internal/instituto-terra"
)

func main() {
	// Routes
	http.HandleFunc("/v1.0/endpoints", internal.GetAllPaths)
	http.HandleFunc("/v1.0/ez/routine", ez.Routine)
	http.HandleFunc("/v1.0/instituto-terra/routine", instituto_terra.Routine)

	// Serving
	fmt.Println("Services status working...")
	http.ListenAndServe(":8080", nil)
}
