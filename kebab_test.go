package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseKebab(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("go-e-c-h-o-framework", str_case.Kebab("GoECHOFramework"))
	assert.Exactly("go-echo-framework", str_case.Kebab("GoEchoFramework"))
	assert.Exactly("go-gin-framework", str_case.Kebab("Go Gin Framework"))
	assert.Exactly("go-echo-framework", str_case.Kebab("  go    Echo      Framework   "))
	assert.Exactly("go-fiber-framework-", str_case.Kebab("goFiberFramework-"))
	assert.Exactly("go-gin-framework", str_case.Kebab("go gin Framework"))
	assert.Exactly("go-gin-frame-work", str_case.Kebab("go gin FrameWork"))
	assert.Exactly("foo-bar", str_case.Kebab("foo-bar"))
	assert.Exactly("foo--bar", str_case.Kebab("Foo-Bar"))
	assert.Exactly("foo_-bar", str_case.Kebab("foo_Bar"))
}