package main

import (
	"github.com/quantum-bridge/core/cmd/bridge"
)

// @title Core bridge API
// @version 1.0
// @description Core bridge API is a service that responsible for the communication between blockchains.
//
// @host localhost:8000
// @schemes http
// @BasePath /v1
func main() {
	bridge.Run()
}
