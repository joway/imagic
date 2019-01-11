package util

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

// LogNail print nail log
func LogNail(msg string) {
	fmt.Printf("%s\r", msg)
}

// LogInfo print info level log
func LogInfo(msg string) {
	color.White(fmt.Sprintf("%s\n", msg))
}

// LogError print error level log
func LogError(err error) {
	color.Red(fmt.Sprintf("Error: %s\n", err))
}

// LogFatal print error level log and exit with code -1
func LogFatal(err error) {
	LogError(err)
	os.Exit(1)
}
