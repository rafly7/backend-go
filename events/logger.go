package events

import (
	"log"
	"os"
	"time"
)

func todayFileName() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func LogFile() *os.File {
	filename := todayFileName()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	return f
}
