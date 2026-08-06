package loadbalencer

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

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request){

	lb.Mutex.Lock()

	server:= lb.Servers[lb.Index];

	lb.Index = (lb.Index + 1) % len(lb.Servers);

	lb.Mutex.Unlock()

	proxy.NewProxy(server).ServeHTTP(w,r)

}