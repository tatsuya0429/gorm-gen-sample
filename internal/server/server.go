package server

import "github.com/labstack/echo/v4"

type Server struct {
	svr *echo.Echo
}

func New() *Server {
	return &Server{
		svr: echo.New(),
	}
}

func (s *Server) Start(addr string) {
	s.svr.Logger.Fatal(s.svr.Start(addr))
}
