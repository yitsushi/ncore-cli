package main

const (
	CLI_BLACK   = "\033[30m"
	CLI_RED     = "\033[31m"
	CLI_GREEN   = "\033[32m"
	CLI_YELLOW  = "\033[33m"
	CLI_BLUE    = "\033[34m"
	CLI_MAGENTA = "\033[35m"
	CLI_CYAN    = "\033[36m"
	CLI_WHITE   = "\033[37m"
	CLI_NOCOLOR = "\033[0m"

	MODE_COLOR   = false
	MODE_NOCOLOR = true
)

var noColor bool = false

func GetTerminalColor(color string) string {
	if noColor {
		return ""
	}
	return color
}

func setColorMode(mode bool) {
	noColor = mode
}
