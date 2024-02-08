package str_case_test

import (
	"testing"

	str_case "github.com/mohprilaksono/str-case"
	"github.com/stretchr/testify/assert"
)

func TestStrApa(t *testing.T) {
	assert := assert.New(t)

	assert.Exactly("Thomas and Friends", str_case.Apa("thomas and friends"))
	assert.Exactly("Thomas and Friends", str_case.Apa("THOMAS AND FRIENDS"))
	assert.Exactly("Thomas and Friends", str_case.Apa("Thomas And Friends"))

	assert.Exactly("Back to the Future", str_case.Apa("back to the future"))
	assert.Exactly("Back to the Future", str_case.Apa("BACK TO THE FUTURE"))
	assert.Exactly("Back to the Future", str_case.Apa("Back To The Future"))

	assert.Exactly("This, Then That", str_case.Apa("this, then that"))
	assert.Exactly("This, Then That", str_case.Apa("THIS, THEN THAT"))
	assert.Exactly("This, Then That", str_case.Apa("This, Then That"))

	assert.Exactly("Bond. James Bond.", str_case.Apa("bond. james bond."))
	assert.Exactly("Bond. James Bond.", str_case.Apa("BOND. JAMES BOND."))
	assert.Exactly("Bond. James Bond.", str_case.Apa("Bond. James Bond."))

	assert.Exactly("Self-Report", str_case.Apa("self-report"))
	assert.Exactly("Self-Report", str_case.Apa("Self-report"))
	assert.Exactly("Self-Report", str_case.Apa("SELF-REPORT"))

	assert.Exactly("As the World Turns, So Are the Days of Our Lives", str_case.Apa("as the world turns, so are the days of our lives"))
	assert.Exactly("As the World Turns, So Are the Days of Our Lives", str_case.Apa("AS THE WORLD TURNS, SO ARE THE DAYS OF OUR LIVES"))
	assert.Exactly("As the World Turns, So Are the Days of Our Lives", str_case.Apa("As The World Turns, So Are The Days Of Our Lives"))

	assert.Exactly("To Kill a Mockingbird", str_case.Apa("to kill a mockingbird"))
	assert.Exactly("To Kill a Mockingbird", str_case.Apa("TO KILL A MOCKINGBIRD"))
	assert.Exactly("To Kill a Mockingbird", str_case.Apa("To Kill A Mockingbird"))
}