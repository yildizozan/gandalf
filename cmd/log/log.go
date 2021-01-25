package log

import (
	"log"
)

func Warning(format string, v ...interface{}) {
	log.Printf(format, v)
}

func Error(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

func Log(format string, v ...interface{}) {
	log.Printf(format, v)
}
