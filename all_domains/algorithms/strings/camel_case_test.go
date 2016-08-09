package strings

import (
	"os"
	. "gopkg.in/check.v1"
)

type CamelCaseSuite struct{}

var _ = Suite(&CamelCaseSuite{})

func (s *CamelCaseSuite) TestCase1(c *C) {
	os.Stdin.WriteString("saveChangesInTheEditor")
	os.Stdin.Close()

	output := captureStdout(func(){
		CamelCase()
	})

	c.Assert(output, Equals, "5\n")
}

func (s *CamelCaseSuite) TestCase2(c *C) {

}
