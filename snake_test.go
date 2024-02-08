package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrSnake(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("laravel_p_h_p_framework", str_case.Snake("LaravelPHPFramework"))
	assert.Exactly("laravel_php_framework", str_case.Snake("LaravelPhpFramework"))
	assert.Exactly("laravel_php_framework", str_case.Snake("Laravel Php Framework"))
	assert.Exactly("laravel_php_framework", str_case.Snake("  Laravel    Php      Framework   "))
	assert.Exactly("laravel_php_framework_", str_case.Snake("LaravelPhpFramework_"))
	assert.Exactly("laravel_php_framework", str_case.Snake("laravel php Framework"))
	assert.Exactly("laravel_php_frame_work", str_case.Snake("laravel php FrameWork"))
	assert.Exactly("foo-bar", str_case.Snake("foo-bar"))
	assert.Exactly("foo-_bar", str_case.Snake("Foo-Bar"))
	assert.Exactly("foo__bar", str_case.Snake("foo_Bar"))
}