package main

import (
	"fmt"
	"io"
	"os"
)

func deferReturnSample(fname string) (err error) {
	var f *os.File
	f, err = os.Create(fname)
	if err != nil {
		return fmt.Errorf("File open error: %v\n", err)
	}
	defer func() {
		err = f.Close()
	}()
	io.WriteString(f, "deferのエラーを拾うサンプル")
	return
}
