package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

type GetInput func() *os.File

func CaptureStdout(f func()) string {
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

func WriteInput(data string) *os.File {
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
