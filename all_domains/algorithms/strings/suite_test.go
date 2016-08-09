package strings

// File contains hook to link gocheck with Golang test system
import (
	"bytes"
	. "gopkg.in/check.v1"
	"io"
	"os"
	"testing"
)

func captureStdout(f func(), str ...string) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// Hook up gocheck into the "go test" runner.
func TestStart(t *testing.T) {
	TestingT(t)
}
