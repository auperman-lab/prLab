package main

import (
	"log"
	"net/http"
)

func main() {

	server := NewApiServer(":3000")

	if err := server.HttpStart(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func (s *HttpServer) HttpStart() error {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: r,
	}
	log.Print("Server has started at address :3000")

	return server.ListenAndServe()

}

type HttpServer struct {
	addr string
}

func NewApiServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}
