package bot_building

import (
	"github.com/iveronanomi/hackerrank/utils"
	. "gopkg.in/check.v1"
)

type BotSavesPrincessSuite struct{}

var _ = Suite(&BotSavesPrincessSuite{})

func (s *BotSavesPrincessSuite) TestBotSavesPrincessCase1(c *C) {
	out := utils.Capture(BotSavesPrincess, "3\n---\n-m-\np--")

	c.Assert(out, Equals, "LEFT\nDOWN\n")
}
