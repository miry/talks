// 03_server.go
//
// Usage:
//  SENTRY_DSN=<url> go run 03_server.go
//  curl -v localhost:8080/panic

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// START GIN OMIT
	if err := initSentry(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := gin.Default()

	// import sentrygin "github.com/getsentry/sentry-go/gin"
	r.Use(sentrygin.New(sentrygin.Options{ // HL
		// To buble panic continue to client. Instead it would return 200.
		Repanic: true,  // HL
	}))  // HL
	// END GIN OMIT

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("Could not resolve a domain name example.com")
	})

	r.GET("/error", func(c *gin.Context) {
		err := errors.New("Huston we have a problem")
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	})

	r.Run()
}

func initSentry() error {
	sentryDSN := os.Getenv("SENTRY_DSN")

	options := sentry.ClientOptions{
		Dsn:              sentryDSN,
		DebugWriter:      os.Stderr,
		Debug:            true,
		Environment:      "development",
		AttachStacktrace: true,
	}

	return sentry.Init(options)
}
