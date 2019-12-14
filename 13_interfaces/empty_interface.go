package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello youtube listeners, this is a set"))
		wc.Close()
	}
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("conversion failed")
	}
}

type Writer interface {
	Write([]byte) (int, error) // method takes []byte, returns int and error
}

type Closer interface {
	Close() error
}

// created interface from other interfaces
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
