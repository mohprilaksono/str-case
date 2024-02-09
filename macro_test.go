package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseMacro(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("GO_E_C_H_O_FRAMEWORK", str_case.Macro("GoECHOFramework"))
	assert.Exactly("GO_ECHO_FRAMEWORK", str_case.Macro("GoEchoFramework"))
	assert.Exactly("GO_GIN_FRAMEWORK", str_case.Macro("Go Gin Framework"))
	assert.Exactly("GO_ECHO_FRAMEWORK", str_case.Macro("  go    Echo      Framework   "))
	assert.Exactly("GO_FIBER_FRAMEWORK_", str_case.Macro("goFiberFramework_"))
	assert.Exactly("GO_GIN_FRAMEWORK", str_case.Macro("go gin Framework"))
	assert.Exactly("GO_GIN_FRAME_WORK", str_case.Macro("go gin FrameWork"))
	assert.Exactly("FOO-BAR", str_case.Macro("foo-bar"))
	assert.Exactly("FOO-_BAR", str_case.Macro("Foo-Bar"))
	assert.Exactly("FOO__BAR", str_case.Macro("foo_Bar"))
}