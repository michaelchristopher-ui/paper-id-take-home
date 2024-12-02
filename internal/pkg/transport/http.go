package transport

import (
	"fmt"
	"net/http"
	"os"
	"paperid-entry-task/internal/conf"
	"time"

	"github.com/labstack/echo"
)

type server struct {
	e            *echo.Echo
	port         string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewServer() server {
	e := echo.New()

	cfg := conf.GetConfig()

	return server{
		e:            e,
		port:         cfg.Server.Port,
		readTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		writeTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}
}

func (h server) GetEcho() *echo.Echo {
	return h.e
}

func (h server) StartServer() {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", h.port),
		ReadTimeout:  h.readTimeout,
		WriteTimeout: h.writeTimeout,
	}
	//This can actually be made to run in a goroutine
	if err := h.e.StartServer(s); err != nil && err != http.ErrServerClosed {
		h.e.Logger.Error(err)
		h.e.Logger.Info("Shutting down the server")
		os.Exit(1)
	}
}
