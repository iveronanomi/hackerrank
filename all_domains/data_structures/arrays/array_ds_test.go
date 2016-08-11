package arrays

import (
	"os"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day2OperatorsSuite struct{}

var _ = Suite(&Day2OperatorsSuite{})

func (s *Day2OperatorsSuite) TestCase1(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("4\n1 4 3 2")
		defer in.Close()
		ArrayDS(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "2 3 4 1")
}
