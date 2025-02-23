package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type LoadBalancer struct {
	Servers       []*Server
	lastServedIdx int
	sync.Mutex
}

var lb *LoadBalancer

func NewLoadBalancer(servers []*Server) *LoadBalancer {
	return &LoadBalancer{Servers: servers}
}

func (lb *LoadBalancer) GetHealthyServer() (*Server, error) {
	lb.Lock()
	defer lb.Unlock()
	for i := 0; i < len(lb.Servers); i++ {
		srv := lb.Servers[(lb.lastServedIdx+i+1)%len(lb.Servers)]
		srv.RLock()
		healthy := srv.Health
		srv.RUnlock()
		if healthy {
			lb.lastServedIdx = (lb.lastServedIdx + i + 1) % len(lb.Servers)
			return srv, nil
		}
	}
	return nil, fmt.Errorf("no healthy servers available")

}

func (lb *LoadBalancer) ForwardRequest(w http.ResponseWriter, r *http.Request) {
	srv, err := lb.GetHealthyServer()
	if err != nil {
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	log.Printf("[REQUEST] forwarding request to %s", srv.Name)
	srv.ReverseProxy.ServeHTTP(w, r)
}
