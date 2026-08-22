package healthcheck

import (
	"log"
	"time"
)

func StartMonitor(route string, servers []string, update func([]string)){

	ticker := time.NewTicker(5* time.Second)

	defer ticker.Stop()

	for range ticker.C {

		// Check all backend servers concurrently.
		healthyServers := CheckAll(servers)

		update(healthyServers)

		log.Println("Health chek:",route, healthyServers)
	}
}