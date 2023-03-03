package server

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Server struct {
	fasthttpServer *fasthttp.Server
}

func (s *Server) Run(port string, hendler fasthttp.RequestHandler) error {
	s.fasthttpServer = &fasthttp.Server{
		Handler:      hendler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.fasthttpServer.ListenAndServe(port)
}
