package main

import (
	"github.com/PrzemyslawMorski/cloud-native-golang/service"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
