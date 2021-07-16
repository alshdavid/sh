package yargs

type MultiString []string

func (i *MultiString) String() string {
	return "my string representation"
}

func (i *MultiString) Set(value string) error {
	*i = append(*i, value)
	return nil
}
