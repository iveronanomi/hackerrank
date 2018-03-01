package bot_building

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type BotSavesPrincess2Suite struct{}

var _ = Suite(&BotSavesPrincess2Suite{})

func (s *BotSavesPrincess2Suite) TestBotSavesPrincess2_Round1(c *C) {
	out := utils.Capture(BotSavesPrincess2, "5\n0 2\n--m--\n-----\np----\n-----\n-----\n")

	c.Assert(out, Equals, "LEFT\n")
}
