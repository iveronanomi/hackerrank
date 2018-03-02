package bot_building

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type BotCleanSuite struct{}

var _ = Suite(&BotCleanSuite{})

func (s *BotCleanSuite) TestBotClean_Step1(c *C) {
	out := utils.Capture(BotClean, "0 0\nbd---\n-d---\n---d-\n---d-\n--d-d\n")

	c.Assert(out, Equals, "RIGHT")
}

func (s *BotCleanSuite) TestBotClean_Step2(c *C) {
	out := utils.Capture(BotClean, "0 1\n-d---\n-d---\n---d-\n---d-\n--d-d\n")

	c.Assert(out, Equals, "CLEAN")
}

func (s *BotCleanSuite) TestBotClean_Step3(c *C) {
	out := utils.Capture(BotClean, "0 1\n-b---\n-d---\n---d-\n---d-\n--d-d\n")

	c.Assert(out, Equals, "DOWN")
}

func (s *BotCleanSuite) TestBotClean_Step4(c *C) {
	out := utils.Capture(BotClean, "1 1\n-----\n-d---\n---d-\n---d-\n--d-d\n")

	c.Assert(out, Equals, "CLEAN")
}

func (s *BotCleanSuite) TestBotClean_Step5(c *C) {
	out := utils.Capture(BotClean, "1 1\n-----\n-b---\n---d-\n---d-\n--d-d\n")

	c.Assert(out, Equals, "RIGHT")
}

