package loadbalancer

import (
	"atulya/api-gateway/proxy"
	"net/http"
	"sync"
)

type LoadBalancer struct{
	Servers[] string
	Index int
	Mutex sync.Mutex
}

func New(servers []string) *LoadBalancer{
	return &LoadBalancer{
		Servers: servers,
	}
}

func (lb *LoadBalancer) UpdateServers(servers []string){
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	lb.Servers = servers
	lb.Index = 0
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request){

	lb.Mutex.Lock()

	if len(lb.Servers) == 0 {
		http.Error(w, "No healthy servers available", http.StatusServiceUnavailable)
		lb.Mutex.Unlock()
		return
	}

	server:= lb.Servers[lb.Index];

	lb.Index = (lb.Index + 1) % len(lb.Servers);

	lb.Mutex.Unlock()

	proxy.NewProxy(server).ServeHTTP(w,r)

}