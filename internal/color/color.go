package color

import (
	"github.com/fatih/color"
)

var (
	Red   = color.New(color.FgRed)
	Green = color.New(color.FgGreen)
)

// Error prints an error message in red
func Error(format string, a ...interface{}) string {
	return Red.Sprintf(format, a...)
}

// Success prints a success message in green
func Success(format string, a ...interface{}) string {
	return Green.Sprintf(format, a...)
}

// PrintError prints an error message in red to stderr
func PrintError(err error) {
	if err != nil {
		Red.Fprintf(color.Output, "Error: %v\n", err)
	}
}

// PrintSuccess prints a success message in green
func PrintSuccess(format string, a ...interface{}) {
	Green.Printf(format+"\n", a...)
}
