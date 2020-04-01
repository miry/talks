package main

import (
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

// START OMIT
func main() {
	options := sentry.ClientOptions{ // HL
		Dsn:              "<SENTRY DSN URL>", // HL
		DebugWriter:      os.Stderr,
		Debug:            true,
		Environment:      "development",
		AttachStacktrace: true,
	}

	err := sentry.Init(options) // HL
	if err != nil {
		fmt.Printf("Sentry initialization error: %v", err)
		os.Exit(1)
	}

	sentry.CaptureMessage("Something went wrong") // HL
	sentry.Flush(time.Second * 5)
}

// END OMIT
