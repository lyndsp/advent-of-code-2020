package passwordpolicy_test

import (
	"integers/passwordpolicy"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_extractPolicyAndPassword(t *testing.T) {

	r, err := passwordpolicy.Parse("1-3 a: abcde")

	require.NoError(t, err)
	assert.Equal(t, 1, r.Min)
	assert.Equal(t, 3, r.Max)
	assert.Equal(t, 'a', r.Policy)
	assert.Equal(t, "abcde", r.Password)
}
func Test_invalidFormatGivesError(t *testing.T) {

	_, err := passwordpolicy.Parse("abcde")

	assert.Error(t, err)
}

func Test_passwordMeetsPolicy(t *testing.T) {

	p := passwordpolicy.PasswordPolicy{
		Min:      1,
		Max:      3,
		Policy:   'a',
		Password: "abcde",
	}
	ok := p.MeetsPolicy()

	assert.True(t, ok)
}
