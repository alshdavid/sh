package handler_help

import (
	"forwardports/src/platform/colors"
)

func Handler() {
	colors.PrintLog(
		"CLI Tool that makes it easy to forward ports between an SSH server and your local machine",
		"",
		"forward-ports --in 8080 --out 8080",
	)
}
