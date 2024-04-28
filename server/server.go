package server

import (
	"fmt"
	"squeeze-auction/internal/config"
	"squeeze-auction/internal/logger"
	"squeeze-auction/server/routes"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	e    *echo.Echo
	port string
	log  *logger.Logger
}

func New(cfg *config.Config, handlers *routes.Routes) *Server {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	itemsGroup := e.Group("/items")
	handlers.RegisterItemsRoutes(itemsGroup)

	return &Server{
		e:    e,
		port: cfg.Port,
		log:  handlers.Log,
	}
}

func (s *Server) Start() error {
	s.log.Info("Starting the server...")
	fmt.Printf("port: %s\n", s.port)
	return s.e.Start(s.port)
}
