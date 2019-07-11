package kqio // Import path github.com/ughoavgfhw/libkq/io

import (
	"io"
	"time"
)

type MessageWriter interface {
	// May ignore the message time.
	WriteMessage(*Message) error
}

type MessageStringWriter interface {
	// May ignore the message time.
	WriteMessageString(*MessageString) error
}

type BasicMessageStringWriter struct {
	w io.Writer
}

func NewMessageStringWriter(writer io.Writer) MessageStringWriter {
	if wc, ok := writer.(io.WriteCloser); ok {
		return NewMessageStringWriteCloser(wc)
	}
	return BasicMessageStringWriter{writer}
}

func (mw BasicMessageStringWriter) WriteMessageString(msg *MessageString) error {
	const bufSize = len(time.RFC3339Nano) + 1 // 1 extra for a comma
	var backing [bufSize]byte
	data := msg.Time.AppendFormat(backing[:0], time.RFC3339Nano)
	if len(data) > 0 {
		data = append(data, ',')
		if err := writeLoop(mw.w, data); err != nil {
			return err
		}
	}
	if err := writeLoop(mw.w, msg.Message); err != nil {
		return err
	}
	backing[0] = '\n'
	return writeLoop(mw.w, backing[:1])
}

func writeLoop(w io.Writer, data []byte) error {
	for len(data) > 0 {
		n, err := w.Write(data)
		if err != nil {
			return err
		}
		data = data[n:]
	}
	return nil
}
