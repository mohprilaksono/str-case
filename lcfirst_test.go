package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseLcFirst(t *testing.T) {
	assert := assert.New(t)
	
	assert.Exactly("go", str_case.LcFirst("Go"))
	assert.Exactly("golang", str_case.LcFirst("Golang"))
	assert.Exactly("go Language", str_case.LcFirst("Go Language"))
	assert.Exactly(" go     Language", str_case.LcFirst(" Go     Language"))
	assert.Exactly("   1go   ", str_case.LcFirst("   1Go   "))
}