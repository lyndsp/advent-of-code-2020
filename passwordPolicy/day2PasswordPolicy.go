package passwordpolicy

import (
	"fmt"
	"unicode/utf8"
)

type (
	// PasswordPolicy details
	PasswordPolicy struct {
		Min, Max int
		Policy   rune
		Password string
	}
)

// Parse does x
func Parse(input string) (PasswordPolicy, error) {
	min, max := 0, 0
	policy, password := "", ""
	_, err := fmt.Sscanf(input, "%d-%d %s %s", &min, &max, &policy, &password)

	if err != nil {
		return PasswordPolicy{}, err
	}
	r, _ := utf8.DecodeRuneInString(policy)

	return PasswordPolicy{
		Min:      min,
		Max:      max,
		Policy:   r,
		Password: password,
	}, nil
}

// MeetsPolicy does x
func (p PasswordPolicy) MeetsPolicy() bool {

	policyCount := 0

	for _, c := range p.Password {
		if c == p.Policy {
			policyCount++
		}
	}

	return policyCount >= p.Min && policyCount <= p.Max
}
