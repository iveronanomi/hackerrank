package cracking_the_coding_interview

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type HashTablesRansomNoteSuite struct{}

var _ = Suite(&HashTablesRansomNoteSuite{})

func (s *HashTablesRansomNoteSuite) TestHashTablesRansomNote_Case1(c *C) {
	out := utils.Capture(HashTablesRansomNote, "6 5\ntwo times three is not four\ntwo times two is four")

	c.Assert(out, Equals, "No")
}

func (s *HashTablesRansomNoteSuite) TestHashTablesRansomNote_Case2(c *C) {
	out := utils.Capture(HashTablesRansomNote, "6 4\ngive me one grand today night\ngive one grand today")

	c.Assert(out, Equals, "Yes")
}

func (s *HashTablesRansomNoteSuite) TestHashTablesRansomNote_Case7(c *C) {
	out := utils.Capture(HashTablesRansomNote, "100 50\ndlji eeyfb ox ayw fmphg x ffbkr z qiq vwvt zqgq nmw hlv c looms ffzif wfzx gzmf ez acidl mxwcw cm ichsi p pltu y jn re enujc ztm k s pkv hv om bsvw czy yzm lpli rj rm waqyk hkxf uffv rctam gp c enbez ngrc uos wfzx z hv acidl mxwcw hlv dlji enujc mxwcw cm p om pkv om x ox enbez pkv enujc rm ngrc x uos dlji zqgq c z rm eeyfb bsvw c dlji gzmf looms re p mxwcw gzmf hlv hv enbez y pkv rm y p gzmf ngrc gzmf wfzx\n\nichsi c c uffv cm jn uffv s rm om uos czy czy nmw hv om gzmf ox uos pltu qiq czy rj vwvt s c ox bsvw acidl ffbkr ez looms ngrc yzm rj dlji wfzx waqyk om looms z om waqyk zqgq wfzx om dlji z nmw mxwcw")

	c.Assert(out, Equals, "No")
}

func (s *HashTablesRansomNoteSuite) TestHashTablesRansomNote_Case4(c *C) {
	out := utils.Capture(HashTablesRansomNote, "12 2\nh ghq g xxy wdnr anjst xxy wdnr h h anjst wdnr\nh ghq")

	c.Assert(out, Equals, "Yes")
}
