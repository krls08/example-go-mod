package go_logger_ex

import "fmt"

var Version string = "1.0"

func Log(msg string) {
	fmt.Printf("\033[01;32m[LOG] %s\033[0m\n", msg)
}
