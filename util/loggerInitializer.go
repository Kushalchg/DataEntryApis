package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

func InitializeLogger() *log.Logger {
	fileName := fmt.Sprintf("logfiles/info_%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		log.Fatal("error while creating log files")
	}

	logger := log.New(file, "info:", log.Ltime|log.Lshortfile)
	return logger
}
