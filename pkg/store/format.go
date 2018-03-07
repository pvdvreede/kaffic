package store

import (
	"bytes"
	"io"

	"github.com/ugorji/go/codec"

	"github.com/pvdvreede/kaffic/pkg/core"
)

const (
	FileFormatVersion = "v1"
)

type MessageFileFormat struct {
	Offset    int64
	Timestamp int64
	Payload   []byte
	CRC       []byte
	Version   string
}

func ToFileReader(msg *MessageFileFormat) (io.Reader, error) {
	h := new(codec.MsgpackHandle)
	w := new(bytes.Buffer)
	enc := codec.NewEncoder(w, h)
	err := enc.Encode(msg)
	return w, err
}

func ToMessageFileFormat(msg *core.InboundMessage) *MessageFileFormat {
	return &MessageFileFormat{
		Offset:    msg.Offset.ToInt64(),
		Timestamp: msg.Timestamp.UnixNano(),
		Version:   FileFormatVersion,
		Payload:   msg.Payload,
	}
}
