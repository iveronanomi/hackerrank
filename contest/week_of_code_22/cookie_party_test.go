package week_of_code_22

import (
	"fmt"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type CookiePartySuite struct{}

var _ = Suite(&CookiePartySuite{})

func (s *CookiePartySuite) TestCases(c *C) {

	testCases := [][3]uint32{
		{3, 2, 1},
		{4, 2, 2},
		{2, 2, 0},
		{2, 5, 1},
		{100, 5, 95},
		{3, 100, 2},
		{10, 1024, 6},
		{0, 0, 0},
		{0, 4, 0},
		{4, 0, 4},
	}

	for _, t := range testCases {
		out := utils.Capture(CookieParty, fmt.Sprintf("%d %d", t[0], t[1]))

		c.Assert(out, Equals, fmt.Sprintf("%d\n", t[2]))
		//c.Logf("Guests: %d, Cookies: %d, Expected: %d = Got: %s", t[0], t[1], t[2], out)
	}
}
