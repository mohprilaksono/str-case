package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrCaseTrain(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("Go-E-C-H-O-Framework", str_case.Train("GoECHOFramework"))
	assert.Exactly("Go-Echo-Framework", str_case.Train("GoEchoFramework"))
	assert.Exactly("Go-Gin-Framework", str_case.Train("go gin framework"))
	assert.Exactly("Go-Echo-Framework", str_case.Train("  go    Echo      Framework   "))
	assert.Exactly("Go-Fiber-Framework-", str_case.Train("goFiberFramework-"))
	assert.Exactly("Go-Gin-Framework", str_case.Train("go gin framework"))
	assert.Exactly("Go-Gin-Frame-Work", str_case.Train("go gin FrameWork"))
	assert.Exactly("Foo-Bar", str_case.Train("foo-bar"))
	assert.Exactly("Foo--Bar", str_case.Train("foo-Bar"))
	assert.Exactly("Foo_-Bar", str_case.Train("foo_Bar"))
	assert.Exactly("The-Train-Case-Converts-Text-To-Train-Case.", str_case.Train("the train case converts text to train Case."))
}