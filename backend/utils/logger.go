// Package utils – application-wide logger instance.
package utils

import (
	"log"
	"os"
)

// Logger is the shared logger for the entire application.
// It writes to stdout with a timestamp prefix.
var Logger = log.New(os.Stdout, "[SUPPLY-CHAIN] ", log.LstdFlags|log.Lmsgprefix)
