package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseAda(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("Go_E_C_H_O_Framework", str_case.Ada("GoECHOFramework"))
	assert.Exactly("Go_Echo_Framework", str_case.Ada("GoEchoFramework"))
	assert.Exactly("Go_Gin_Framework", str_case.Ada("Go Gin Framework"))
	assert.Exactly("Go_Echo_Framework", str_case.Ada("  go    Echo      Framework   "))
	assert.Exactly("Go_Fiber_Framework_", str_case.Ada("goFiberFramework_"))
	assert.Exactly("Go_Gin_Framework", str_case.Ada("go gin Framework"))
	assert.Exactly("Go_Gin_Frame_Work", str_case.Ada("go gin FrameWork"))
	assert.Exactly("Foo-Bar", str_case.Ada("foo-bar"))
	assert.Exactly("Foo-_Bar", str_case.Ada("Foo-Bar"))
	assert.Exactly("Foo__Bar", str_case.Ada("foo_Bar"))
}