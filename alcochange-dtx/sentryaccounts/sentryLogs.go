package sentryaccounts

import (
	sentry "github.com/getsentry/sentry-go"
)

func SentryLog(message string) {
	sentry.CaptureMessage(message)
}

func SentryLogExceptions(err error) {
	sentry.CaptureException(err)
}
