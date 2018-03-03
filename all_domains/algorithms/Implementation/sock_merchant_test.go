package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type SockMerchantSuite struct{}

var _ = Suite(&SockMerchantSuite{})

func (s *SockMerchantSuite) TestSockMerchant_Case1(c *C) {
	out := utils.Capture(SockMerchant, "9\n10 20 20 10 10 30 50 10 20")

	c.Assert(out, Equals, "3")
}
