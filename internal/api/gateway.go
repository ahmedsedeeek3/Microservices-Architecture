package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"MicroserviceArchitecture/internal/discovery"
)

type APIGateway struct {
	serviceDiscovery *discovery.ServiceDiscovery
}

func NewAPIGateway(sd *discovery.ServiceDiscovery) *APIGateway {
	return &APIGateway{serviceDiscovery: sd}
}

func (g *APIGateway) HandleRequest(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 2 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	serviceName := pathParts[1]

	serviceEndpoint, ok := g.serviceDiscovery.Discover(serviceName)
	if !ok {
		http.Error(w, fmt.Sprintf("Service '%s' not found", serviceName), http.StatusNotFound)
		return
	}

	targetURL := serviceEndpoint + r.URL.Path
	proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	for header, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Error forwarding request to microservice", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for header, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(header, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("Error copying response: %v", err)
	}
}
