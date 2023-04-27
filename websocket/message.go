package websocket

import (
	"github.com/gorilla/websocket"
	"time"
)

type TextMessage struct {
	msg           string
	sourceAddress string
	targetAddress string
	time          time.Time
}

func NewSendMessage(conn *websocket.Conn, msg string) TextMessage {
	return TextMessage{
		msg:           msg,
		sourceAddress: conn.LocalAddr().String(),
		targetAddress: conn.RemoteAddr().String(),
		time:          time.Now(),
	}
}

func NewReceiveMessage(conn *websocket.Conn, msg string) TextMessage {
	return TextMessage{
		msg:           msg,
		sourceAddress: conn.RemoteAddr().String(),
		targetAddress: conn.LocalAddr().String(),
		time:          time.Now(),
	}
}

func (h TextMessage) Message() string {
	return h.msg
}

func (h TextMessage) TimeString() string {
	return h.time.Format("2006-01-02 15:04:05")
}

func (h TextMessage) Time() time.Time {
	return h.time
}

func (h TextMessage) TargetAddress() string {
	return h.targetAddress
}

func (h TextMessage) SourceAddress() string {
	return h.sourceAddress
}

func (h TextMessage) ToString() string {
	return "[" + h.sourceAddress + " --> " + h.targetAddress + "] " + h.time.Format("2006-01-02 15:04:05") + " : " + h.msg
}
