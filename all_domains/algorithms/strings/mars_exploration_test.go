package strings

import (
	. "gopkg.in/check.v1"
	"os"
)

type MarsExplorationSuite struct{}

var _ = Suite(&MarsExplorationSuite{})

func (s *MarsExplorationSuite) TestCase1(c *C) {
	out := captureStdout(func() {
		in := writeInput("SOSSPSSQSSOR")
		defer in.Close()
		MarsExploration(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "3\n")
}

func (s *MarsExplorationSuite) TestCase2(c *C) {
	out := captureStdout(func() {
		in := writeInput("SOSSOT")
		defer in.Close()
		MarsExploration(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "1\n")
}
