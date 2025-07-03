package commons

import (
	"io"
	"log"
	"os"
)

var (
	Logger        = log.New(os.Stdout, "info  ", log.LstdFlags)
	VerboseLogger = log.New(io.Discard, "debug ", log.LstdFlags)
)

func EnableVerboseLogging() {
	VerboseLogger.SetOutput(os.Stdout)
}
