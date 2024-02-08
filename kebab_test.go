package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrKebab(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("laravel-p-h-p-framework", str_case.Kebab("LaravelPHPFramework"))
	assert.Exactly("laravel-php-framework", str_case.Kebab("LaravelPhpFramework"))
	assert.Exactly("laravel-php-framework", str_case.Kebab("Laravel Php Framework"))
	assert.Exactly("laravel-php-framework", str_case.Kebab("  Laravel    Php      Framework   "))
	assert.Exactly("laravel-php-framework-", str_case.Kebab("LaravelPhpFramework-"))
	assert.Exactly("laravel-php-framework", str_case.Kebab("laravel php Framework"))
	assert.Exactly("laravel-php-frame-work", str_case.Kebab("laravel php FrameWork"))
	assert.Exactly("foo-bar", str_case.Kebab("foo-bar"))
	assert.Exactly("foo--bar", str_case.Kebab("Foo-Bar"))
	assert.Exactly("foo_-bar", str_case.Kebab("foo_Bar"))
}