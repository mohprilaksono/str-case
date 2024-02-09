package str_case_test

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrCaseSponge(t *testing.T) {
	value := "HTML is not a programming language"
	isUpper := false
	isLower := false
	for range value {
		if rand.Float32() > 0.5 {
			isUpper = true
		} else {
			isLower = true
		}
	}

	assert := assert.New(t)

	assert.True(isUpper)
	assert.True(isLower)
}