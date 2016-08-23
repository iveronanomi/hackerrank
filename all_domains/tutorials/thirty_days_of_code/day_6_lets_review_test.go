package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day6LetsReviewSuite struct{}

var _ = Suite(&Day6LetsReviewSuite{})

func (s *Day6LetsReviewSuite) TestCase1(c *C) {
	out := utils.CaptureOut(Day6LetsReview, "2\nHacker\nRank")

	c.Assert(out, Equals, "Hce akr\nRn ak\n")
}
