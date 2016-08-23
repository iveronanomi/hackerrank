package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day8DictionariesAndMapsSuite struct{}

var _ = Suite(&Day8DictionariesAndMapsSuite{})

func (s *Day8DictionariesAndMapsSuite) TestCase1(c *C) {
	out := utils.CaptureOut(DictionariesAndMaps, "3\nsam 99912222\ntom 11122222\nharry 12299933\nsam\nedward\nharry")

	c.Assert(out, Equals, "sam=99912222\nNot found\nharry=12299933\n")
}
