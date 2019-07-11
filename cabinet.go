package libkq

import (
	"errors"
	"log"
	"time"

	"github.com/ughoavgfhw/libkq/io"
	"github.com/ughoavgfhw/libkq/parser"
)

// Cabinet manages a connection to a killer queen cabinet. It reads and parses
// messages from the contained reader, and encodes and writes to the contained
// writer. The writer may be nil, in which case sending messages will fail.
// Whenever an "alive" message is read, a reply will automatically be sent if
// the writer is available.
type Cabinet struct {
	Source kqio.MessageStringReader
	Reply  kqio.MessageStringWriter
	Parser parser.Parser
}

// Creates a cabinet with the given reader. If the reader is a
// kqio.MessageStringReaderWriter, it will also be used as the reply writer.
// Otherwise there will be no reply writer.
func NewCabinet(source kqio.MessageStringReader) *Cabinet {
	reply, _ := source.(kqio.MessageStringWriter)
	return &Cabinet{
		Source: source,
		Reply: reply,
		Parser: parser.Parser{},
	}
}

func (cab *Cabinet) ReadMessage(out* kqio.Message) error {
	var msg kqio.MessageString
	if err := cab.Source.ReadMessageString(&msg); err != nil { return err }
	key, val, err := cab.Parser.Parse(msg.Message)
	if err != nil { return err }
	out.Time = msg.Time
	out.Type = kqio.MessageType(string(key))
	out.Val = val

	if out.Type == "alive" && cab.Reply != nil {
		reply := kqio.MessageString{
			Time: time.Now(),
			Message: []byte("![k[im alive],v[null]]!"),
		}
		if err := cab.Reply.WriteMessageString(&reply); err != nil {
			log.Print("kq cabinet: \"alive\" reply failed: ", err.Error())
		}
	}
	return nil
}

func (cab *Cabinet) WriteMessage(msg* kqio.Message) error {
	return errors.New("libkq.Cabinet.WriteMessage: unimplemented")
}
