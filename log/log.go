package log

import (
	"log"
	"os"
)

var (
	Error   = log.New(os.Stderr, "", 0)
	Warn    = log.New(os.Stderr, "", 0)
	Default = log.New(os.Stdout, "", 0)
)
