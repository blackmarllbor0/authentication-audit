package server

import (
	"aptekaaprel/config/configGetter"
	"aptekaaprel/internal/pkg/signal"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server

	configGetter configGetter.ConfigGetter
}

func NewServer(configGetter configGetter.ConfigGetter) *Server {
	return &Server{
		configGetter: configGetter,
	}
}

func (s *Server) Run(handlers http.Handler) error {
	s.server = &http.Server{
		Addr:           fmt.Sprintf(":%v", s.configGetter.GetValueByKey("SERVER_PORT")),
		Handler:        handlers,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	go signal.ListenSignals()

	return s.server.ListenAndServe()
}
