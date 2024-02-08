package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrStudly(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("LaravelPHPFramework", str_case.Studly("laravel_p_h_p_framework"))
	assert.Exactly("LaravelPhpFramework", str_case.Studly("laravel_php_framework"))
	assert.Exactly("LaravelPhPFramework", str_case.Studly("laravel-phP-framework"))
	assert.Exactly("LaravelPhpFramework", str_case.Studly("laravel  -_-  php   -_-   framework   "))
	assert.Exactly("FooBar", str_case.Studly("fooBar"))
	assert.Exactly("FooBar", str_case.Studly("foo_bar"))
	assert.Exactly("FooBar", str_case.Studly("foo_bar"))
	assert.Exactly("FooBarBaz", str_case.Studly("foo-barBaz"))
	assert.Exactly("FooBarBaz", str_case.Studly("foo-bar_baz"))
}