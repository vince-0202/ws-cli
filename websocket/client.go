package websocket

import (
	"github.com/desertbit/grumble"
	"github.com/gorilla/websocket"
)

type Client struct {
	serverPath            string
	connection            *websocket.Conn
	app                   *grumble.App
	sendMessageHistory    []TextMessage
	receiveMessageHistory []TextMessage
}

func NewWsClient(serverPath string, app *grumble.App) *Client {
	return &Client{
		serverPath:            serverPath,
		app:                   app,
		sendMessageHistory:    make([]TextMessage, 0, 10),
		receiveMessageHistory: make([]TextMessage, 0, 10),
	}
}

func (c *Client) ServerPath() string {
	return c.serverPath
}

func (c *Client) Connection() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.serverPath, nil)
	c.connection = conn
	if err != nil {
		return err
	}
	go c.HandlerTextReceive()
	return nil
}

func (c *Client) CloseConnection() error {
	return c.connection.Close()
}

func (c *Client) SendText(msg string) error {
	c.sendMessageHistory = append(c.sendMessageHistory, NewSendMessage(c.connection, msg))
	return c.connection.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (c *Client) HandlerTextReceive() {
	for {
		messageType, message, err := c.connection.ReadMessage()
		if err != nil {
			break
		}
		if messageType == websocket.TextMessage {
			msg := NewReceiveMessage(c.connection, string(message))
			c.receiveMessageHistory = append(c.receiveMessageHistory, msg)
			c.app.Println("------- receive ", msg.TimeString(), " -------")
			c.app.Println(msg.Message())
			c.app.Println("---------------------------------------------")
		}
	}
}

func (c *Client) HistoryOfSend() []TextMessage {
	return c.sendMessageHistory
}
func (c *Client) HistoryOfReceive() []TextMessage {
	return c.receiveMessageHistory
}
