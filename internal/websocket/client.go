package websocket

import (
	"github.com/desertbit/grumble"
	"github.com/gorilla/websocket"
)

// Client  websocket client
type Client struct {
	serverPath            string
	connection            *websocket.Conn
	app                   *grumble.App
	sendMessageHistory    []TextMessage
	receiveMessageHistory []TextMessage
}

// NewWsClient the websocket client constructor
func NewWsClient(serverPath string, app *grumble.App) *Client {
	return &Client{
		serverPath:            serverPath,
		app:                   app,
		sendMessageHistory:    make([]TextMessage, 0, 10),
		receiveMessageHistory: make([]TextMessage, 0, 10),
	}
}

// ServerPath return server path which is the client to connection
func (c *Client) ServerPath() string {
	return c.serverPath
}

// Connection to websocket server
func (c *Client) Connection() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.serverPath, nil)
	c.connection = conn
	if err != nil {
		return err
	}
	go c.HandlerTextReceive()
	return nil
}

// CloseConnection close the websocket connection
func (c *Client) CloseConnection() error {
	return c.connection.Close()
}

// SendText send a text message to server
func (c *Client) SendText(msg string) error {
	c.sendMessageHistory = append(c.sendMessageHistory, NewSendMessage(c.connection, msg))
	return c.connection.WriteMessage(websocket.TextMessage, []byte(msg))
}

// HandlerTextReceive handler the text message from websocket server
// to print on cli
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

// HistoryOfSend the history of the send messages
func (c *Client) HistoryOfSend() []TextMessage {
	return c.sendMessageHistory
}

// HistoryOfReceive the history of the received messages
func (c *Client) HistoryOfReceive() []TextMessage {
	return c.receiveMessageHistory
}
