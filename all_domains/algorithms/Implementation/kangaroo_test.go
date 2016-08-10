package implementation

import (
	"os"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type KangarooSuite struct{}

var _ = Suite(&KangarooSuite{})

func (s *KangarooSuite) TestCase1(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("0 3 4 2")
		defer in.Close()
		Kangaroo(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "YES\n")
}

func (s *KangarooSuite) TestCase2(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("0 2 5 3")
		defer in.Close()
		Kangaroo(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "NO\n")
}
