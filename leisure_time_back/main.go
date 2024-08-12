package main

import (
	"leisure_time_back/cmd/server"
)

func main() {
	server := server.NewServer()

	server.SetupServer()

	server.RunServer()
}
