package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	servers := []*Server{
		NewServer("server-1", "http://127.0.0.1:5001"),
		NewServer("server-2", "http://127.0.0.1:5002"),
		NewServer("server-3", "http://127.0.0.1:5003"),
		NewServer("server-4", "http://127.0.0.1:5004"),
		NewServer("server-5", "http://127.0.0.1:5005"),
	}

	lb = NewLoadBalancer(servers)

	go startHealthCheck(lb)

	http.HandleFunc("/", lb.ForwardRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("[STARTING] Load balancer running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
