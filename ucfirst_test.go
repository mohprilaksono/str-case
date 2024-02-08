package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrUcFirst(t *testing.T) {
	assert := assert.New(t)
	
	assert.Exactly("Go", str_case.UcFirst("go"))
	assert.Exactly("Golang", str_case.UcFirst("golang"))
	assert.Exactly("Go Language", str_case.UcFirst("go Language"))
	assert.Exactly(" Go     Language", str_case.UcFirst(" go     Language"))
	assert.Exactly("   Go   ", str_case.UcFirst("   go   "))
}