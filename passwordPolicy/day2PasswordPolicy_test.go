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

	testData := []struct {
		input       string
		meetsPolicy bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", true},
	}

	for _, td := range testData {
		t.Run(td.input, func(t *testing.T) {
			p, err := passwordpolicy.Parse(td.input)
			require.NoError(t, err)

			ok := p.MeetsPolicy()

			assert.Equal(t, td.meetsPolicy, ok)
		})
	}
}

func Test_day2_answer(t *testing.T) {
	okCount := 0

	for _, v := range day2TestInput {
		p, err := passwordpolicy.Parse(v)
		require.NoError(t, err)

		if p.MeetsPolicy() {
			okCount++
		}
	}

	assert.Equal(t, 0, okCount)
}
