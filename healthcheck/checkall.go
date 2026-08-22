package healthcheck

type Result struct {
	ServerName string
	Healthy    bool
}

func CheckAll(servers []string) []string {

	results := make(chan Result, len(servers))

	// Start one goroutine for each backend.
	for _, server := range servers {

		go func(server string) {

			// Recover from an unexpected panic.
			defer func() {
				if r := recover(); r != nil {
					results <- Result{
						ServerName: server,
						Healthy:    false,
					}
				}
			}()

			// Check the health of this backend.
			healthy := IsHealthy(server)

			// Send the result back through the channel.
			results <- Result{
				ServerName: server,
				Healthy:    healthy,
			}

		}(server)
	}

	healthyServers := []string{}

	// Receive one result for every server we checked.
	for i := 0; i < len(servers); i++ {

		result := <-results

		if result.Healthy {
			healthyServers = append(
				healthyServers,
				result.ServerName,
			)
		}
	}

	return healthyServers
}