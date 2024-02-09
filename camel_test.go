package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseCamel(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("goECHOFramework", str_case.Camel("Go_e_c_h_o_framework"))
	assert.Exactly("goGinFramework", str_case.Camel("Go_gin_framework"))
	assert.Exactly("goEchOFramework", str_case.Camel("Go-echO-framework"))
	assert.Exactly("goFiberFramework", str_case.Camel("Go  -_-  fiber   -_-   framework   "))
	assert.Exactly("fooBar", str_case.Camel("FooBar"))
	assert.Exactly("fooBar", str_case.Camel("Foo_bar"))
	assert.Exactly("fooBar", str_case.Camel("Foo_bar"))
	assert.Exactly("fooBarBaz", str_case.Camel("Foo-barBaz"))
	assert.Exactly("fooBarBaz", str_case.Camel("Foo-bar_baz"))
}