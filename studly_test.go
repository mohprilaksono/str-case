package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrStudly(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("GoGINFramework", str_case.Studly("Go_g_i_n_framework"))
	assert.Exactly("GoFiberFramework", str_case.Studly("Go_fiber_framework"))
	assert.Exactly("GoEchOFramework", str_case.Studly("Go-echO-framework"))
	assert.Exactly("GoFiberFramework", str_case.Studly("Go  -_-  fiber   -_-   framework   "))
	assert.Exactly("FooBar", str_case.Studly("fooBar"))
	assert.Exactly("FooBar", str_case.Studly("foo_bar"))
	assert.Exactly("FooBar", str_case.Studly("foo_bar"))
	assert.Exactly("FooBarBaz", str_case.Studly("foo-barBaz"))
	assert.Exactly("FooBarBaz", str_case.Studly("foo-bar_baz"))
}