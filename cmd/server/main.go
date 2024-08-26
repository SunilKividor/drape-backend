package main

import (
	"log"

	"github.com/SunilKividor/drape/internal/server"
)

func main() {
	s := server.GetServer(":3000")
	err := s.StartServer()
	if err != nil {
		log.Fatalf("error: Error starting the server")
	}
}
