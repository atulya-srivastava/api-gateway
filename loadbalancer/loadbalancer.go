package loadbalancer

import (
	"atulya/api-gateway/proxy"
	"net/http"
	"sync"
)

type LoadBalancer struct {
	servers []string
	index   int
	mu      sync.Mutex
	proxies map[string]http.Handler
}

func New(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		proxies: make(map[string]http.Handler),
	}
}

func (lb *LoadBalancer) UpdateServers(servers []string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.servers = servers

	// Keep the index valid after the server list changes.
	if len(lb.servers) == 0 {
		lb.index = 0
		return
	}

	lb.index = lb.index % len(lb.servers)
}


func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request){

	lb.mu.Lock()

	if len(lb.servers) == 0 {
		http.Error(w, "No healthy servers available", http.StatusServiceUnavailable)
		lb.mu.Unlock()
		return
	}

	targetServer := lb.servers[lb.index]
	lb.index = (lb.index + 1) % len(lb.servers)

	targetProxy, exists := lb.proxies[targetServer]
	if !exists {
		targetProxy = proxy.NewProxy(targetServer)
		lb.proxies[targetServer] = targetProxy
	}

	lb.mu.Unlock()

	targetProxy.ServeHTTP(w, r)
}