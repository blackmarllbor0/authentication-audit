package server

import (
	"auth_audit/config/configValueGetter"
	"auth_audit/pkg/signal"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server

	configValueGetter configValueGetter.ConfigValueGetter
}

func NewServer(configValueGetter configValueGetter.ConfigValueGetter) *Server {
	return &Server{
		configValueGetter: configValueGetter,
	}
}

func (s *Server) Run(handlers http.Handler) error {
	s.server = &http.Server{
		Addr:           fmt.Sprintf(":%v", s.configValueGetter.GetValueByKeys("app.server.port")),
		Handler:        handlers,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	go signal.ListenSignals()

	return s.server.ListenAndServe()
}
