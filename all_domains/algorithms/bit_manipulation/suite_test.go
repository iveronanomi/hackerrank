package bit_manipulation

// File contains hook to link gocheck with Golang test system
import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestStart(t *testing.T) {
	TestingT(t)
}
