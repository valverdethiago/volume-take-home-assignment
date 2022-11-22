package server

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/valverdethiago/volume-take-home-assignment/config"
	"net/http"
	"time"
)

type Controller interface {
	ConfigureRoutes(router *gin.Engine)
}

type Server struct {
	router *gin.Engine
	config *config.AppConfig
}

// NewServer creates a new server instance
func NewServer(config config.AppConfig) *Server {
	server := &Server{
		router: gin.Default(),
		config: &config,
	}
	server.ConfigureLogging()
	return server
}

func (server *Server) ConfigureController(controller Controller) {
	controller.ConfigureRoutes(server.router)
}

// ConfigureLogging configure gin logs
func (server *Server) ConfigureLogging() {
	flag.Parse()
	server.router.Use(ginglog.Logger(3 * time.Second))
	server.router.Use(gin.Recovery())
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start() error {
	readTimeout := time.Duration(server.config.ReadTimeout) * time.Second
	writeTimeout := time.Duration(server.config.WriteTimeout) * time.Second
	serverAddress := fmt.Sprintf("0.0.0.0:%s", server.config.Port)

	s := &http.Server{
		Addr:         serverAddress,
		Handler:      server.router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return s.ListenAndServe()
}
