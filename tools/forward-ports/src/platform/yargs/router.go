package yargs

import (
	"fmt"
	"forwardports/src/platform/colors"
)

type RouterHandler func(args []string)

type Router struct {
	commands map[string]RouterHandler
	command  string
	args     []string
}

func NewRouter(args []string) *Router {
	command, args := newArgs(args)
	return &Router{
		commands: map[string]RouterHandler{},
		command:  command,
		args:     args,
	}
}

func (r *Router) Command(command string, handler RouterHandler) {
	r.commands[command] = handler
}

func (r *Router) Exec() {
	found := r.commands[r.command]

	if found == nil {
		colors.PrintError("Command not recognized")
		colors.PrintLog("Commands:")
		for command := range r.commands {
			fmt.Println("  ", command)
		}
		return
	}

	found(r.args)
}
