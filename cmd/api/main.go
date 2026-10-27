package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	// Keep this scaffold local and route-free until API implementation begins.
	if err := router.Run(net.JoinHostPort("127.0.0.1", port)); err != nil {
		log.Fatal(err)
	}
}
