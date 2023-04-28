package websocket

import (
	"github.com/desertbit/grumble"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

// ServerType the type of websocket server
type ServerType int

const (
	// Normal normal websocket server
	Normal ServerType = iota
	// Echo echo server
	Echo
)

// ServerTypeEnum the type of websocket server
var ServerTypeEnum = []string{"normal", "echo"}

// ToString return the type of websocket server
func (st ServerType) ToString() string {
	return ServerTypeEnum[st]
}

// ToServerType string to ServerType
func ToServerType(t string) ServerType {
	if t == ServerTypeEnum[Echo] {
		return Echo
	}
	return Normal
}

// Server websocket socket server
type Server struct {
	port                  int
	serverPath            string
	serverType            ServerType
	connections           map[string]*websocket.Conn
	app                   *grumble.App
	sendMessageHistory    []TextMessage
	receiveMessageHistory []TextMessage
	wsUpgrader            websocket.Upgrader
}

// NewWsServer the constructor of Server
func NewWsServer(path string, port int, serverType ServerType, app *grumble.App) *Server {
	if path[0] != '/' {
		path = "/" + path
	}
	return &Server{
		port:                  port,
		serverType:            serverType,
		serverPath:            path,
		app:                   app,
		connections:           make(map[string]*websocket.Conn, 10),
		sendMessageHistory:    make([]TextMessage, 0, 10),
		receiveMessageHistory: make([]TextMessage, 0, 10),
		wsUpgrader:            websocket.Upgrader{},
	}
}

// Run start the server
func (s *Server) Run(stop chan struct{}) {
	http.HandleFunc(s.serverPath, s.handleWebSocket)
	if err := http.ListenAndServe(":"+strconv.Itoa(s.port), nil); err != nil {
		stop <- struct{}{}
	}
}

// SendTextToAllConnection send text message to all connection
func (s *Server) SendTextToAllConnection(msg string) error {
	for _, conn := range s.connections {
		s.sendMessageHistory = append(s.sendMessageHistory, NewSendMessage(conn, msg))
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			return err
		}
	}
	return nil
}

// SendText send text message to a connection
func (s *Server) SendText(remoteAddr string, msg string) error {
	if conn, ok := s.connections[remoteAddr]; ok {
		s.sendMessageHistory = append(s.sendMessageHistory, NewSendMessage(conn, msg))
		return conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
	return nil
}

// handleWebSocket handle the message
func (s *Server) handleWebSocket(writer http.ResponseWriter, request *http.Request) {
	s.wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	upgradeConnection, err := s.wsUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}
	s.connections[upgradeConnection.RemoteAddr().String()] = upgradeConnection
	for {
		if upgradeConnection == nil {
			break
		}
		messageType, message, err := upgradeConnection.ReadMessage()
		if err != nil {
			break
		}
		if messageType == websocket.TextMessage {
			msg := NewReceiveMessage(upgradeConnection, string(message))
			s.receiveMessageHistory = append(s.receiveMessageHistory, msg)
			s.app.Println("")
			s.app.Println("------- receive ", msg.TimeString(), " -------")
			s.app.Println(msg.Message())
			s.app.Println("---------------------------------------------")
			s.app.Println("")
			if s.serverType == Echo {
				if err = s.SendText(upgradeConnection.RemoteAddr().String(), msg.Message()); err != nil {
					break
				}
			}
		}
	}
}

// ServerPath return the server path
func (s *Server) ServerPath() string {
	return s.serverPath
}

// Clients return the list of all client
func (s *Server) Clients() []string {
	res := make([]string, 0, len(s.connections))
	for remoteAddr := range s.connections {
		res = append(res, remoteAddr)
	}
	return res
}

// Port return the server port
func (s *Server) Port() int {
	return s.port
}

// URL return the url
func (s *Server) URL() string {
	return "ws://127.0.0.1:" + strconv.Itoa(s.Port()) + s.ServerPath()
}

// HistoryOfSend the history of the send messages
func (s *Server) HistoryOfSend() []TextMessage {
	return s.sendMessageHistory
}

// HistoryOfReceive the history of the received messages
func (s *Server) HistoryOfReceive() []TextMessage {
	return s.receiveMessageHistory
}
