package go_logger_ex

import (
	"log"
)

var Version string = "2.0"

type Logger interface {
	Info(string)
	//Debug(string)
	Warning(string)
	Error(string)
}

type logger struct {
}

func NewLogger() Logger {
	return &logger{}
}

func (l logger) Info(msg string) {
	log.Printf("\033[01;32m[INFO]\t %s\033[0m\n", msg)
}

func (l logger) Error(msg string) {
	log.Printf("\033[01;33m[WARNING]\t %s\033[0m\n", msg)
}

func (l logger) Warning(msg string) {
	log.Printf("\033[01;33m[DEBUG]\t %s\033[0m\n", msg)
}
