// 05_server.go
//
// Usage:
//  SENTRY_DSN=<url> go run 05_server.go
//  curl -v localhost:8080/error

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("Could not resolve a domain name example.com")
	})

	l := initLogger()
	// START OMIT
	r.GET("/error", func(c *gin.Context) {
		err := processRequest() // HL
		err = fmt.Errorf("Access Denied: %w", err)  // HL
		l.Warn().Err(err).Str("endpoint", "error_log").Msg("Unbelivable")
		c.JSON(200, gin.H{"message": "warning"})
	})
	// END OMIT

	r.Run()
}

func initSentry() error {
	sentryDSN := os.Getenv("SENTRY_DSN")

	options := sentry.ClientOptions{
		Dsn:              sentryDSN, // HL
		DebugWriter:      os.Stderr,
		Debug:            true,
		Environment:      "development",
		AttachStacktrace: true,
	}

	return sentry.Init(options)
}

func initLogger() zerolog.Logger {
	zerolog.ErrorMarshalFunc = func(err error) interface{} {
		sentry.CaptureException(err)
		return err
	}

	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func processRequest() error {
	return errors.New("Huston we have a problem")
}
