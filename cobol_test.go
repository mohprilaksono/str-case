package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseCobol(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("GO-E-C-H-O-FRAMEWORK", str_case.Cobol("GoECHOFramework"))
	assert.Exactly("GO-ECHO-FRAMEWORK", str_case.Cobol("GoEchoFramework"))
	assert.Exactly("GO-GIN-FRAMEWORK", str_case.Cobol("Go Gin Framework"))
	assert.Exactly("GO-ECHO-FRAMEWORK", str_case.Cobol("  go    Echo      Framework   "))
	assert.Exactly("GO-FIBER-FRAMEWORK-", str_case.Cobol("goFiberFramework-"))
	assert.Exactly("GO-GIN-FRAMEWORK", str_case.Cobol("go gin Framework"))
	assert.Exactly("GO-GIN-FRAME-WORK", str_case.Cobol("go gin FrameWork"))
	assert.Exactly("FOO-BAR", str_case.Cobol("foo-bar"))
	assert.Exactly("FOO--BAR", str_case.Cobol("Foo-Bar"))
	assert.Exactly("FOO_-BAR", str_case.Cobol("foo_Bar"))
}