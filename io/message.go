package kqio // Import path github.com/ughoavgfhw/libkq/io

import "time"

type MessageType string
type Message struct {
	Time time.Time
	Type MessageType
	Val  interface{}
}

type MessageString struct {
	Time    time.Time
	Message []byte
}
