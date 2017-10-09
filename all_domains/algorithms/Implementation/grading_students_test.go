package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type GradingStudentsSuite struct{}

var _ = Suite(&GradingStudentsSuite{})

func (s *GradingStudentsSuite) TestCase1(c *C) {
	out := utils.CaptureOut(GradingStudents, "4\n73\n67\n38\n33")

	c.Assert(out, Equals, "75\n67\n40\n33\n")
}
