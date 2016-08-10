package strings

// File contains hook to link gocheck with Golang test system
import (
	"bytes"
	. "gopkg.in/check.v1"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func captureStdout(f func()) string {
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

func writeInput(data string) *os.File {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(in, data)
	if err != nil {
		panic(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}
	return in
}

// Hook up gocheck into the "go test" runner.
func TestStart(t *testing.T) {
	TestingT(t)
}
