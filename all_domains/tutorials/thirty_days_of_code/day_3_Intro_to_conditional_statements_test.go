package thirty_days_of_code

import (
	"os"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day3IntroTtoConditionalStatementsSuite struct{}

var _ = Suite(&Day3IntroTtoConditionalStatementsSuite{})

func (s *Day3IntroTtoConditionalStatementsSuite) TestCase1(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("3")
		defer in.Close()
		Day3IntroTtoConditionalStatements(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "Weird")
}

func (s *Day3IntroTtoConditionalStatementsSuite) TestCase2(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("24")
		defer in.Close()
		Day3IntroTtoConditionalStatements(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "Not Weird")
}
