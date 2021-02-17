package main

import (
	"fmt"

	s "github.com/CharlesTenorio/servicos/server"
)

func main() {

	server, err := s.New()
	if err != nil {
		fmt.Errorf("error on initializing the server: %v", err)
		return
	}

	if err := server.Start(); err != nil {
		fmt.Errorf("error on server start: %v", err)
	}
}
