package main

import (
	"log"
	"net/http"

	"MicroserviceArchitecture/internal/api"
	"MicroserviceArchitecture/internal/discovery"
	"MicroserviceArchitecture/internal/microservice"
)

func main() {
	// Initialize components-discovery part
	sd := discovery.NewServiceDiscovery()
	sd.Register("service1", "http://localhost:8081")
	sd.Register("service2", "http://localhost:8082")

	gateway := api.NewAPIGateway(sd)
	http.HandleFunc("/", gateway.HandleRequest)

	// Start microservices
	go microservice.Start("service1", ":8081")
	go microservice.Start("service2", ":8082")

	// Start API Gateway
	log.Println("Starting API Gateway on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
