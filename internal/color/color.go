package color

import (
	"github.com/fatih/color"
)

var (
	Red   = color.New(color.FgRed)
	Green = color.New(color.FgGreen)
)

// PrintError prints an error message in red
func PrintError(err error) {
	if err != nil {
		Red.Fprintf(color.Output, "Error: %v\n", err)
	}
}

// PrintSuccess prints a success message in green
func PrintSuccess(format string, a ...interface{}) {
	Green.Printf(format+"\n", a...)
}
