package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day4ClassVsInstanceSuite struct{}

var _ = Suite(&Day4ClassVsInstanceSuite{})

func (s *Day4ClassVsInstanceSuite) TestCase1(c *C) {
	out := utils.CaptureOut(Day4ClassVsInstance, "4\n-1\n10\n16\n18")

	c.Assert(out, Equals,
		"Age is not valid, setting age to 0.\n"+
			"You are young.\n"+
			"You are young.\n\n"+
			"You are young.\n"+
			"You are a teenager.\n\n"+
			"You are a teenager.\n"+
			"You are old.\n\n"+
			"You are old.\n"+
			"You are old.\n\n",
	)
}
