package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type ClimbingTheLeaderboardSuite struct{}

var _ = Suite(&ClimbingTheLeaderboardSuite{})

func (s *ClimbingTheLeaderboardSuite) TestClimbingTheLeaderboardCase1(c *C) {
	out := utils.CaptureOut(ClimbingTheLeaderboard, "7\n100 100 50 40 40 20 10\n4\n5 25 50 120")

	c.Assert(out, Equals, "6\n4\n2\n1\n")
}
