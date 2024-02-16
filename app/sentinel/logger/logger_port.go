package logger

import "os"

func SentinelLoggerUrl() string {
	u := os.Getenv("SENTINEL_LOGGER_URL")
	if u == "" {
		return "[::]:50051"
	}
	return u
}
