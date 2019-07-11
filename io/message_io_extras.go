package kqio // Import path github.com/ughoavgfhw/libkq/io

import "io"

type MessageReaderWriter interface {
	MessageReader
	MessageWriter
}

type MessageStringReadWriter interface {
	MessageStringReader
	MessageStringWriter
}

type MessageStringReadCloser interface {
	MessageStringReader
	io.Closer
}

type MessageStringWriteCloser interface {
	MessageStringWriter
	io.Closer
}

type MessageStringReadWriteCloser interface {
	MessageStringReader
	MessageStringWriter
	io.Closer
}

type BasicMessageStringReadCloser struct {
	BasicMessageStringReader
	io.Closer
}

type BasicMessageStringWriteCloser struct {
	BasicMessageStringWriter
	// Closer extracted from the writer.
}

func NewMessageStringReadCloser(rc io.ReadCloser) BasicMessageStringReadCloser {
	return NewMessageStringReader(rc).(BasicMessageStringReadCloser)
}

func NewMessageStringWriteCloser(wc io.WriteCloser) BasicMessageStringWriteCloser {
	return BasicMessageStringWriteCloser{BasicMessageStringWriter{wc}}
}

func (mwc BasicMessageStringWriteCloser) Close() error {
	return mwc.w.(io.Closer).Close()
}
