package sentryaccounts

import (
	"ecargoware/alcochange-dtx/conf"
	"fmt"
	"time"

	sentry "github.com/getsentry/sentry-go"
)

func InitiateSentryLog() {
	var sentryURL = "https://f33ec3afbd084974871dde304e760777@o458167.ingest.sentry.io/5455220"
	if conf.ServerFlag {
		sentryURL = "https://1a0dc07beb7f458c87ab7b3c947db2c7@o501960.ingest.sentry.io/5583758"
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryURL,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
}
