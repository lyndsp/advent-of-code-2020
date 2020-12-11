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

type (
	// NewPasswordPolicy details
	NewPasswordPolicy struct {
		Index1, Index2 int
		Policy         string
		Password       string
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

// ParseNew does x
func ParseNew(input string) (NewPasswordPolicy, error) {
	position1, position2 := 0, 0
	policy, password := "", ""
	_, err := fmt.Sscanf(input, "%d-%d %s %s", &position1, &position2, &policy, &password)

	if err != nil {
		return NewPasswordPolicy{}, err
	}

	return NewPasswordPolicy{
		Index1:   position1 - 1,
		Index2:   position2 - 1,
		Policy:   string(policy[0]),
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

// MeetsNewPolicy does x
func (p NewPasswordPolicy) MeetsNewPolicy() bool {

	sAtFirstIndex := string(p.Password[p.Index1])
	sAtSecondIndex := string(p.Password[p.Index2])

	if sAtFirstIndex == p.Policy && sAtSecondIndex != p.Policy {
		return true
	}

	if sAtFirstIndex != p.Policy && sAtSecondIndex == p.Policy {
		return true
	}

	return false
}
