package microservice

import (
	"fmt"
	"log"
	"net/http"
)

type Microservice struct {
	Name string
}

func (m *Microservice) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Microservice %s handling request", m.Name)
}

func Start(name, addr string) {
	ms := &Microservice{Name: name}
	http.HandleFunc("/"+name, ms.Handle)
	log.Printf("Starting %s on %s\n", name, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
