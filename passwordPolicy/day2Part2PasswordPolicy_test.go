package passwordpolicy_test

import (
	"integers/passwordpolicy"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_extractNewPolicyAndPassword(t *testing.T) {

	r, err := passwordpolicy.ParseNew("1-3 a: abcde")

	require.NoError(t, err)
	assert.Equal(t, 0, r.Index1)
	assert.Equal(t, 2, r.Index2)
	assert.Equal(t, "a", r.Policy)
	assert.Equal(t, "abcde", r.Password)
}

func Test_passwordMeetsNewPolicy(t *testing.T) {

	testData := []struct {
		input       string
		meetsPolicy bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", false},
	}

	for _, td := range testData {
		t.Run(td.input, func(t *testing.T) {
			p, err := passwordpolicy.ParseNew(td.input)
			require.NoError(t, err)

			ok := p.MeetsNewPolicy()

			assert.Equal(t, td.meetsPolicy, ok)
		})
	}
}

func Test_day2_part2_answer(t *testing.T) {
	okCount := 0

	for _, v := range day2TestInput {
		p, err := passwordpolicy.ParseNew(v)
		require.NoError(t, err)

		if p.MeetsNewPolicy() {
			okCount++
		}
	}

	assert.Equal(t, 705, okCount)
}
