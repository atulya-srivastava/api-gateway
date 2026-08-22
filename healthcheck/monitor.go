package healthcheck

import (
	"log"
	"time"
	"atulya/api-gateway/loadbalancer"
)

func StartMonitor(servers [] string,lb *loadbalancer.LoadBalancer) {

	ticker := time.NewTicker(5* time.Second)

	defer ticker.Stop()

	for range ticker.C {

		// Check all backend servers concurrently.
		healthyServers := CheckAll(servers)

		lb.UpdateServers(healthyServers)

		log.Println("Healthy servers:", healthyServers)
	}
}