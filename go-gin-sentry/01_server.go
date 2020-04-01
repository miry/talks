// 01_server.go
//
// Usage:
//  go run 01_server.go
//  curl -v localhost:8080/ping

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
