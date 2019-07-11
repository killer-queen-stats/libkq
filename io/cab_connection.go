package kqio // Import path github.com/ughoavgfhw/libkq/io

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type CabConnection struct {
	ws *websocket.Conn
}

func Connect(cabUrlString string) (*CabConnection, error) {
	return ConnectWithDialer(websocket.DefaultDialer, cabUrlString)
}
func ConnectWithDialer(dialer *websocket.Dialer, cabUrlString string) (*CabConnection, error) {
	ws, _, err := dialer.Dial(cabUrlString, nil)
	if err != nil {
		return nil, err
	}
	return &CabConnection{ws}, nil
}

func (conn *CabConnection) ReadMessageString(out *MessageString) error {
	_, buf, err := conn.ws.ReadMessage()
	out.Time = time.Now()
	out.Message = buf

	closeErr, ok := err.(*websocket.CloseError)
	switch {
	case !ok:
		return err
	case closeErr.Code == websocket.CloseNormalClosure:
		return io.EOF
	default:
		return &ConnectionLostError{closeErr}
	}
}

func (conn *CabConnection) WaitForClose() error {
	var err error
	for err == nil {
		var buff []byte
		_, buff, err = conn.ws.ReadMessage()
		if buff != nil {
			if len(buff) > 128 {
				buff = buff[:128]
			}
			log.Printf("kq cab connection: dropping unexpected message: %q", string(buff))
		}
	}
	closeErr, ok := err.(*websocket.CloseError)
	switch {
	case !ok:
		return err
	case closeErr.Code == websocket.CloseNormalClosure:
		return nil
	default:
		return &ConnectionLostError{closeErr}
	}
}

func (conn *CabConnection) WriteMessageString(msg *MessageString) error {
	return conn.ws.WriteMessage(websocket.TextMessage, msg.Message)
}

const closeTimeLimit = time.Second

func (conn *CabConnection) Close() error {
	deadline := time.Now().Add(closeTimeLimit)
	sendErr := conn.ws.WriteControl(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "connection closed by client"),
		deadline)
	conn.ws.SetReadDeadline(deadline)
	waitErr := conn.WaitForClose()
	closeErr := conn.ws.Close()
	if sendErr != nil {
		return sendErr
	}
	if waitErr != nil {
		return waitErr
	}
	return closeErr
}

type ConnectionLostError struct {
	Err *websocket.CloseError
}

func (e *ConnectionLostError) Error() string {
	return fmt.Sprint("kq cab connection lost: ", e.Err.Error())
}
