package main

import (
	"forwardports/src/platform/yargs"
	"os"
)

var Routes = struct {
	Empty   string
	Version string
	Help    string
}{
	Empty:   "",
	Version: "version",
	Help:    "help",
}

func main() {
	router := yargs.NewRouter(os.Args[1:])

	router.Command(Routes.Version, version.Handler())
	router.Command(Routes.Serve, serve.Handler())

	router.Exec()
}
