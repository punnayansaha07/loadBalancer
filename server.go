package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
	sync.RWMutex
}

func NewServer(name, urlStr string) *Server {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("Invalid server URL %s: %v", urlStr, err)
	}
	return &Server{
		Name:         name,
		URL:          urlStr,
		ReverseProxy: httputil.NewSingleHostReverseProxy(u),
		Health:       true,
	}
}

func (s *Server) CheckHealth() {
	resp, err := http.Head(s.URL)
	status := err == nil && resp.StatusCode == http.StatusOK

	s.Lock()
	s.Health = status
	s.Unlock()

	if status {
		log.Printf("[HEALTH CHECK] ✅ %s is healthy", s.Name)
	} else {
		log.Printf("[HEALTH CHECK] ❌ %s is unhealthy", s.Name)
	}
}
