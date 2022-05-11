package main

import (
	"MatchingAPI/server"
)

// @title Matching API
// @version 1.0
// @description Matching API to send request Driver Location API

// @host localhost:3001
func main() {
	s := server.NewServer()
	s.StartServer(3001)
}
