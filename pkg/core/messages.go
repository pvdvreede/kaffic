package core

import (
	"time"
)

type InboundMessage struct {
	Offset    Offset
	Payload   []byte
	Timestamp time.Time
}

func NewInboundMessage(os Offset, t time.Time, payload []byte) *InboundMessage {
	return &InboundMessage{
		Offset:    os,
		Timestamp: t,
		Payload:   payload,
	}
}
