package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseSwap(t *testing.T) {
	assert := assert.New(t)
	
	assert.Exactly("gO", str_case.Swap("Go"))
	assert.Exactly("gOLANg", str_case.Swap("GolanG"))
	assert.Exactly("gO lANGUAGE", str_case.Swap("Go Language"))
	assert.Exactly(" gO     LANGUAGe", str_case.Swap(" Go     languagE"))
	assert.Exactly("   1go   ", str_case.Swap("   1GO   "))
}