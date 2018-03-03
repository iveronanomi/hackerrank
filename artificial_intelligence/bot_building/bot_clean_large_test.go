package bot_building

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type BotCleanLargeSuite struct{}

var _ = Suite(&BotCleanLargeSuite{})

func (s *BotCleanLargeSuite) TestBotCleanLarge_Round1(c *C) {
	out := utils.Capture(BotSavesPrincess2, "0 0\n5 5\nb---d\n-d--d\n--dd-\n--d--\n----d\n")

	c.Assert(out, Equals, "RIGHT\n")
}
