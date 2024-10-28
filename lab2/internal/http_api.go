package internal

import (
	"log"
	"net/http"
)

type HttpServer struct {
	addr string
}

func NewApiServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (s *HttpServer) HttpStart() error {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: r,
	}
	log.Print("Server has started at address :3000")

	return server.ListenAndServe()

}
