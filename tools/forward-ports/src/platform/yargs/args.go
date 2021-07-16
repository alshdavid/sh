package yargs

func newArgs(a []string) (command string, args []string) {
	if len(a) >= 1 {
		command = a[0]
	}
	if len(a) >= 2 {
		args = append(args, a[1:]...)
	}

	return command, args
}
