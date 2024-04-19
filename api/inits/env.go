package inits

import (
	"errors"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/jasonpatricklee/gowebserver/utils"
	"github.com/joho/godotenv"
)

func ReadEnv() {
	err := godotenv.Load()
	utils.HandleError(err)
}
func StartSentry() {
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		utils.HandleError(errors.New("sentry dsn is not set"))
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		EnableTracing:    true,
		TracesSampleRate: 0.30,
	}); err != nil {
		utils.HandleError(err)
	}
}
