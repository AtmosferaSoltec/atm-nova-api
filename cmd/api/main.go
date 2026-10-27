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

	// Loopback by default. The container image sets HOST=0.0.0.0 so Docker can
	// publish the port; the host nginx stays the only public entry point.
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	// Keep this scaffold route-free until API implementation begins.
	if err := router.Run(net.JoinHostPort(host, port)); err != nil {
		log.Fatal(err)
	}
}
