package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type GradingStudentsSuite struct{}

var _ = Suite(&GradingStudentsSuite{})

func (s *GradingStudentsSuite) TestGradingStudents_Case1(c *C) {
	out := utils.Capture(GradingStudents, "4\n73\n67\n38\n33")

	c.Assert(out, Equals, "75\n67\n40\n33\n")
}
