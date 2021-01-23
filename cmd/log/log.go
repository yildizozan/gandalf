package log

import "log"

func Warning(message string) {
	log.Printf("[Warning]\t%s\n", message)
}

func Error(message string) {
	log.Fatalf("[Error]\t%s\n", message)
}

func Log(message string) {
	log.Printf("[Log]\t%s", message)
}
