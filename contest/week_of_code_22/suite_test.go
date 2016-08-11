package week_of_code_22

// File contains hook to link gocheck with Golang test system
import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestStart(t *testing.T) {
	TestingT(t)
}
