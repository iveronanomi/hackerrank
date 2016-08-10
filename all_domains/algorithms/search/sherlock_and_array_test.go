package search

import (
	"os"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type SherlockAndArraySuite struct{}

var _ = Suite(&SherlockAndArraySuite{})

func (s *SherlockAndArraySuite) TestCase1(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("2\n3\n1 2 3\n4\n1 2 3 3")
		defer in.Close()
		SherlockAndArray(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "NO\nYES\n")
}
