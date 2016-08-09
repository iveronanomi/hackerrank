package strings

// File contains hook to link gocheck with Golang test system
import (
	"testing"
	"os"
	"bytes"
	"io"
	. "gopkg.in/check.v1"
)
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// Hook up gocheck into the "go test" runner.
func TestStart(t *testing.T) {
	TestingT(t)
}
