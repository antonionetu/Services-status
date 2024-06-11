package pkg

import (
	"net/http"
	"service-status/config"
)

func GetStatus(project string, endpoint string) int {
	client := &http.Client{
		Timeout: config.RequestTimeout,
	}

	fullUrl := project + endpoint

	resp, err := client.Get(fullUrl)
	if err != nil {
		return 408
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
