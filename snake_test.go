package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseSnake(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("go_e_c_h_o_framework", str_case.Snake("GoECHOFramework"))
	assert.Exactly("go_echo_framework", str_case.Snake("GoEchoFramework"))
	assert.Exactly("go_gin_framework", str_case.Snake("Go Gin Framework"))
	assert.Exactly("go_echo_framework", str_case.Snake("  go    Echo      Framework   "))
	assert.Exactly("go_fiber_framework_", str_case.Snake("goFiberFramework_"))
	assert.Exactly("go_gin_framework", str_case.Snake("go gin Framework"))
	assert.Exactly("go_gin_frame_work", str_case.Snake("go gin FrameWork"))
	assert.Exactly("foo-bar", str_case.Snake("foo-bar"))
	assert.Exactly("foo-_bar", str_case.Snake("Foo-Bar"))
	assert.Exactly("foo__bar", str_case.Snake("foo_Bar"))
}