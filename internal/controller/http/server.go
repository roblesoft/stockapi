package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/roblesoft/stockapi/internal/helper"
)

type Server struct {
	Port   int
	Router *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Port: 3005,
	}
}

func (s *Server) SetupRouter() {
	s.Router = http.NewServeMux()

	s.Router.HandleFunc("GET /ping", s.Ping)
}

func (s Server) Run() {
	log.Print("Starting server")
	stack := helper.CreateStack(
		helper.Loggin,
	)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: stack(s.Router),
	}
	log.Fatal(server.ListenAndServe())
}
