package server

import (
	"github.com/SunilKividor/drape/internal/router"
	"github.com/gin-gonic/gin"
)

type server struct {
	port string
}

func GetServer(port string) server {
	return server{
		port: port,
	}
}

func (s *server) StartServer() error {
	engine := gin.Default()
	router.Router(engine)
	return engine.Run(s.port)
}
