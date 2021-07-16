package yargs

type Params struct{}

func Parse(input string) *Params {
	// values := map[string][]string{}
	// has := map[string]bool{}

	// fields := strings.Fields(input)

	return &Params{}
}

func (p *Params) Has(key string) bool {
	return true
}

func (p *Params) Values(key string) []string {
	return []string{}
}
