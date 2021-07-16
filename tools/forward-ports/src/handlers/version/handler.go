package handler_version

import (
	"forwardports/src/platform/colors"
	"forwardports/src/platform/info"
)

func Handler() {
	colors.PrintLog(
		"Random Generation Tool",
		"Version: "+info.Version,
	)
}
