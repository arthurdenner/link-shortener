package utils

import (
	"fmt"
	"log"
)

// Logger prints a given string if the `l` flag is true
func Logger(isLogsOn *bool, format string, values ...interface{}) {
	if *isLogsOn {
		log.Printf(fmt.Sprintf("%s\n", format), values...)
	}
}
