package log

import "github.com/fatih/color"

var colorLog *color.Color
var logLevel Level

func init() {
	colorLog = color.New()
}
