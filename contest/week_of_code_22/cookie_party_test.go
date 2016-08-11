package week_of_code_22

import (
	"fmt"
	"os"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type CookiePartySuite struct{}

var _ = Suite(&CookiePartySuite{})

func (s *CookiePartySuite) TestCases(c *C) {

	cases := [][3]uint32{{3, 2, 1}, {4, 2, 2}, {2, 2, 0}, {2, 5, 1}, {100, 5, 95}, {3, 100, 2}, {10, 1024, 6}, {0, 0, 0}, {0, 4, 0}, {4, 0, 4}}

	for _, cas := range cases {
		out := utils.CaptureStdout(func() {
			in := utils.WriteInput(fmt.Sprintf("%d %d", cas[0], cas[1]))
			defer in.Close()
			CookieParty(func() *os.File {
				return in
			})
		})

		c.Assert(out, Equals, fmt.Sprintf("%d\n", cas[2]))
		//c.Logf("Guests: %d, Cookies: %d, Expected: %d = Got: %s", cas[0], cas[1], cas[2], out)
	}
}
