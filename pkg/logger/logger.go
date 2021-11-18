package logger

import (
	"log"
	"os"
)

type LoggerInstance struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

func NewLogger() *LoggerInstance {

	InfoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &LoggerInstance{InfoLogger: InfoLogger, WarningLogger: WarningLogger, ErrorLogger: ErrorLogger}

}
