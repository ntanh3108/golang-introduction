package main

import (
	"bytes"
	"fmt"
)

func main() {
	var writer Writer = &ConsoleWriter{}
	writer.Write([]byte("Hello World"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 1; i <= 5; i++ {
		fmt.Println(inc.Increment())
	}

	var buffered WriterCloser = NewBufferedWriterCloser()
	buffered.NewWrite([]byte("Hello lai la minh Nguyen Tuan Anh day!!!"))
	buffered.Close()

}

type Writer interface {
	Write([]byte) (int, error)
}

type NewWriter interface {
	NewWrite([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	NewWriter
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) NewWrite(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err = bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, err
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

type Incrementer interface {
	Increment() int
}

type ConsoleWriter struct{}

type IntCounter int

func (cv ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
