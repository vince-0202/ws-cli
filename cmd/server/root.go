package server

import (
	"github.com/desertbit/grumble"
	"github.com/vince-0202/ws-cli/websocket"
)

var (
	wsServer *websocket.Server
	stop     chan struct{}
)

var serverCmd = &grumble.Command{
	Name:    "server",
	Aliases: []string{"s"},
	Help:    "websocket server",
	Args: func(a *grumble.Args) {
		a.String("path", "server path")
	},
	Flags: func(f *grumble.Flags) {
		f.Int("p", "port", 8443, "server port")
		f.String("t", "type", "normal", "server type")
	},
	Run: func(c *grumble.Context) error {
		wsServer = websocket.NewWsServer(c.Args.String("path"), c.Flags.Int("port"), websocket.ToServerType(c.Flags.String("type")), serverApp)
		go wsServer.Run(stop)
		if err := serverApp.Run(); err != nil {
			return err
		}
		return nil
	},
}

var serverApp = grumble.New(&grumble.Config{
	Name:        "websocket server",
	Description: "websocket testing tools",
	Flags: func(f *grumble.Flags) {
	},
})

func init() {
	serverApp.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("                                             ")
		a.Println("__      _____   ___  ___ _ ____   _____ _ __ ")
		a.Println("\\ \\ /\\ / / __| / __|/ _ \\ '__\\ \\ / / _ \\ '__|")
		a.Println(" \\ V  V /\\__ \\ \\__ \\  __/ |   \\ V /  __/ |   ")
		a.Println("  \\_/\\_/ |___/ |___/\\___|_|    \\_/ \\___|_|   ")
		a.Println(" =======================================================  ")
		a.Println("  websocket server : ", wsServer.URL())
		a.Println(" =======================================================  ")
		a.Println("                                             ")
	})
	serverApp.AddCommand(clientList)
	serverApp.AddCommand(send)
	serverApp.AddCommand(sendHistory)
	serverApp.AddCommand(receiveHistory)

}

// Register register a children grumble.App to root grumble.App
func Register(root *grumble.App) {
	root.AddCommand(serverCmd)
	serverApp.OnClose(func() error {
		if err := root.Close(); err != nil {
			return err
		}
		return nil
	})
}
