package logger

import (
	"fmt"
	"strings"
	"time"
)

var currentTime = func() string {
	return time.Now().Local().Format(time.DateTime)
}

func LogTextBuilder(a ...any) string {
	logPrefix := fmt.Sprintf("[LOG][%s]", currentTime())
	var messageParts []string

	for _, element := range a {
		messageParts = append(messageParts, fmt.Sprintf("%v", element))
	}

	message := strings.Join(messageParts, " ")
	finalLog := fmt.Sprintf("%s %s\n", logPrefix, message)

	return finalLog
}
