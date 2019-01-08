package util

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func LogInfo(msg string) {
	color.White(fmt.Sprintf("imagic: %s\n", msg))
}

func LogError(err error) {
	color.Red(fmt.Sprintf("Error: %s\n", err))
}

func LogFatal(err error) {
	LogError(err)
	os.Exit(1)
}
