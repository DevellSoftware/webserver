package server

import "net/http"

type Server struct {
}

func NewServer() Server {
	s := Server{}
	return s
}

func (s *Server) Listen(port int) {
  mux := http.NewServeMux()
  mux.Han
