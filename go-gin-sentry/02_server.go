// 02_server.go
//
// Usage:
//  go run 02_server.go
//  curl -v localhost:8080/panic
//  curl -v localhost:8080/error

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/panic", func(c *gin.Context) { 									// HL
		panic("Could not resolve a domain name example.com")	// HL
	})																											// HL

	r.GET("/error", func(c *gin.Context) { 												// HL
		err := errors.New("Huston we have a problem") 							// HL
		c.Error(err)																								// HL
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	// HL
	})																														// HL

	r.Run()
}
