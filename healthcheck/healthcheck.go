package healthcheck

import (
	"net/http"
	"time"
)

var client =http.Client{
	Timeout: 2* time.Second,
}

func IsHealthy(server string) bool{

	resp,err := client.Get(server +"/health");

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK

}