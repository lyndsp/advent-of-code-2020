package integers

import "strings"

func policyAndPassword(input string) (string, string) {
	components := strings.Split(input, ": ")

	if len(components) != 2 {
		return "", ""
	}

	return components[0], components[1]
}
