package util

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

// LogInfo print info level log
func LogInfo(msg string) {
	color.White(fmt.Sprintf("imagic: %s\n", msg))
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
