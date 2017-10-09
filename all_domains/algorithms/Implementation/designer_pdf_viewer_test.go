package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type DesignerPDFViewerSuite struct{}

var _ = Suite(&DesignerPDFViewerSuite{})

func (s *DesignerPDFViewerSuite) TestDesignerPDFViewerCase1(c *C) {
	out := utils.CaptureOut(DesignerPDFViewer, "1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5\nabc")

	c.Assert(out, Equals, "9\n")
}

func (s *DesignerPDFViewerSuite) TestDesignerPDFViewerCase2(c *C) {
	out := utils.CaptureOut(DesignerPDFViewer, "1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 7\nzaba")

	c.Assert(out, Equals, "28\n")
}
