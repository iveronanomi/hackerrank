package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

func Capture(execute func(), inputData string) string {
	//Create input, and write into it
	os.Stdin = setData(inputData)
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	//Execute function with passed input
	execute()

	os.Stdin.Close()
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func setData(data string) (f *os.File) {
	var err error
	if f, err = ioutil.TempFile("", ""); err != nil {
		panic(err)
	}
	if _, err = io.WriteString(f, data); err != nil {
		panic(err)
	}
	if _, err = f.Seek(0, io.SeekStart); err != nil {
		panic(err)
	}
	return f
}
