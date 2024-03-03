package main

import (
	"fmt"
	"regexp"
)

func main() {
	logMessage := "[INFO][2024-02-21 22:27:47] VSECMSENTINEL Test message"
	regexString := `^\[INFO\]\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \w+ Test message$`

	matched, err := regexp.MatchString(regexString, logMessage)
	if err != nil {
		println("Error:", err)
		return
	}

	if matched {
		println("The log message matches the regular expression.")
	} else {
		println("The log message does not match the regular expression.")
	}
}
