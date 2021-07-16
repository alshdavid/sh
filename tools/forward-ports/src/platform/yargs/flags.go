package yargs

import (
	"strings"
)

type FlagCollection struct {
	properties map[string][]string
}

func NewFlagCollection(input string) *FlagCollection {
	properties := map[string][]string{}

	flags := &FlagCollection{
		properties: properties,
	}

	fields := strings.Fields(input)

	for i, field := range fields {
		if strings.HasPrefix(field, "--") {
			if properties[field[2:]] == nil {
				properties[field[2:]] = []string{}
			}
			continue
		}
		if strings.HasPrefix(field, "-") {
			for _, r := range field {
				char := string(r)
				properties[char] = []string{"true"}

			}
			continue
		}
		properties[fields[i-1][2:]] = append(properties[fields[i-1][2:]], field)
	}

	for key, fields := range properties {
		if len(fields) == 0 {
			properties[key] = append(properties[key], "true")
		}
	}

	return flags
}

func (f *FlagCollection) Get(flag string) []string {
	found := f.properties[flag]
	if found == nil {
		return []string{}
	}
	return found
}

func (f *FlagCollection) Has(flag string) bool {
	return f.properties[flag] != nil
}
