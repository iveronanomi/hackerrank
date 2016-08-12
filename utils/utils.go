package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

func CaptureOut(f func(), inputData string) string {
	//Create input, and write into it
	in := writeInput(inputData)
	os.Stdin = in
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	//Execute function with passed input
	f()

	in.Close()
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
