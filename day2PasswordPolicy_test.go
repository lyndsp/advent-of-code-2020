package integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractPolicyAndPassword(t *testing.T) {

	policy, password := policyAndPassword("1-3 a: abcde")

	assert.Equal(t, "1-3 a", policy)
	assert.Equal(t, "abcde", password)
}
