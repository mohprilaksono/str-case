package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseTitle(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("Jefferson Costella  ", str_case.Title("jeFFersOn CostellA  "))
	assert.Exactly("Jefferson   Costella", str_case.Title("JEFfersOn   costElla"))
}