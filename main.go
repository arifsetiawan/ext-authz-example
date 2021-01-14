package main

import (
	"log"
	"net/http"
)

// Server implements a HTTP handler that validates a request body.
type Server struct {
}

// Check request object.
func (s *Server) check(w http.ResponseWriter, r *http.Request) {
	log.Println("#############################")
	log.Println("new requests ... ")
	log.Println("[method] ", r.Method)
	log.Println("[uri] ", r.RequestURI)
	// Print all headers
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			log.Println("[header] ", name, ":", value)
		}
	}

	// All is OK
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	log.Print("serving on port 8088")
	server := &Server{}

	http.HandleFunc("/extauth", server.check)

	http.HandleFunc("/", server.check)

	//log.Fatal(http.ListenAndServe(":8088", nil))

	err := http.ListenAndServeTLS(":8088", "certs/ext-authz.crt", "certs/ext-authz.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
