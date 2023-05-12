package websocket

import (
	"github.com/gorilla/websocket"
	"time"
)

// TextMessage the text-message object
type TextMessage struct {
	msg           string
	sourceAddress string
	targetAddress string
	time          time.Time
}

// NewSendMessage TextMessage constructor for send message
func NewSendMessage(conn *websocket.Conn, msg string) TextMessage {
	return TextMessage{
		msg:           msg,
		sourceAddress: conn.LocalAddr().String(),
		targetAddress: conn.RemoteAddr().String(),
		time:          time.Now(),
	}
}

// NewReceiveMessage TextMessage constructor for receive message
func NewReceiveMessage(conn *websocket.Conn, msg string) TextMessage {
	return TextMessage{
		msg:           msg,
		sourceAddress: conn.RemoteAddr().String(),
		targetAddress: conn.LocalAddr().String(),
		time:          time.Now(),
	}
}

// Message return the context of message
func (h TextMessage) Message() string {
	return h.msg
}

// TimeString return the time of message
func (h TextMessage) TimeString() string {
	return h.time.Format("2006-01-02 15:04:05")
}

// Time return the time of message
func (h TextMessage) Time() time.Time {
	return h.time
}

// TargetAddress return the target address of message
func (h TextMessage) TargetAddress() string {
	return h.targetAddress
}

// SourceAddress return the source address of message
func (h TextMessage) SourceAddress() string {
	return h.sourceAddress
}

// ToString return the message info
func (h TextMessage) ToString() string {
	return "[" + h.sourceAddress + " --> " + h.targetAddress + "] " + h.time.Format("2006-01-02 15:04:05") + " : " + h.msg
}
