package healthcheck

import (
	"net/http"
	"time"
)

func IsHealthy(server string) bool{

	client :=http.Client{
		Timeout: 2* time.Second,
	}

	resp,err := client.Get(server +"/health");

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK

}