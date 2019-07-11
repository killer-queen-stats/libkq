package kqio // Import path github.com/ughoavgfhw/libkq/io

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"time"
)

type MessageReader interface {
	ReadMessage(out *Message) error
}

type MessageStringReader interface {
	// The buffer in the result is owned by the reader. It may replace the
	// message buffer with a slice into an internal buffer, or write into the
	// buffer, extending it as needed. If the implementation ever returns an
	// internal buffer, it must be acceptable but not required to pass the
	// same internal buffer to a future read call. The result is valid until
	// the next read.
	ReadMessageString(out *MessageString) error
}

type InvalidMessageStringError struct {
	Err     string
	Message string
}

func (e *InvalidMessageStringError) Error() string {
	return fmt.Sprintf("invalid message string: %v: %q", e.Err, e.Message)
}

type BasicMessageStringReader struct {
	scanner *bufio.Scanner
}

func NewMessageStringReader(reader io.Reader) MessageStringReader {
	msr := BasicMessageStringReader{bufio.NewScanner(reader)}
	if rc, ok := reader.(io.ReadCloser); ok {
		return BasicMessageStringReadCloser{msr, rc}
	}
	return msr
}

func (r BasicMessageStringReader) ReadMessageString(out *MessageString) error {
	if !r.scanner.Scan() {
		e := r.scanner.Err()
		if e == nil {
			e = io.EOF
		}
		return e
	}
	buf := r.scanner.Bytes()
	if buf[0] == '!' {
		out.Time = time.Now()
		out.Message = buf
		return nil
	}
	return SplitTimeFromMessageString(buf, out)
}

func SplitTimeFromMessageString(buf []byte, out *MessageString) error {
	timeEnd := -1
	for i, b := range buf {
		if b == ',' {
			timeEnd = i
			break
		}
	}
	if timeEnd == -1 {
		return &InvalidMessageStringError{"did not find separator between time and message", string(buf)}
	}
	timeStr := string(buf[:timeEnd])
	if t, err := time.Parse(time.RFC3339Nano, timeStr); err == nil {
		out.Time = t
		out.Message = buf[timeEnd+1:]
		return nil
	}
	if ms, err := strconv.ParseInt(timeStr, 10, 64); err == nil {
		timestampNanos := ms * int64(time.Millisecond/time.Nanosecond)
		out.Time = time.Unix(0, timestampNanos).Local()
		out.Message = buf[timeEnd+1:]
		return nil
	}
	return &InvalidMessageStringError{"could not parse message time", string(buf)}
}
