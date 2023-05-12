package client

import (
	"github.com/desertbit/grumble"
	"github.com/vince-0202/ws-cli/internal/websocket"
)

var (
	wsClient *websocket.Client
)

var clientCmd = &grumble.Command{
	Name:    "client",
	Aliases: []string{"ct"},
	Help:    "websocket client testing tools",
	Args: func(a *grumble.Args) {
		a.String("path", "server path")
	},
	Run: func(c *grumble.Context) error {
		wsClient = websocket.NewWsClient(c.Args.String("path"), clientApp)
		if err := wsClient.Connection(); err != nil {
			return err
		}
		if err := clientApp.Run(); err != nil {
			return err
		}
		return nil
	},
}

var clientApp = grumble.New(&grumble.Config{
	Name:        "websocket client",
	Description: "websocket testing tools",
	Flags: func(f *grumble.Flags) {
	},
})

func init() {
	clientApp.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("                     _ _            _   ")
		a.Println("__      _____    ___| (_) ___ _ __ | |_ ")
		a.Println("\\ \\ /\\ / / __|  / __| | |/ _ \\ '_ \\| __|")
		a.Println(" \\ V  V /\\__ \\ | (__| | |  __/ | | | |_ ")
		a.Println("  \\_/\\_/ |___/  \\___|_|_|\\___|_| |_|\\__|")
		a.Println(" =======================================================  ")
		a.Println("  websocket server : ", wsClient.ServerPath())
		a.Println(" =======================================================  ")
		a.Println("   ")
	})
	clientApp.OnClose(func() error {
		if err := wsClient.CloseConnection(); err != nil {
			return err
		}
		clientApp.Println("Closed the websocket connection: ", wsClient.ServerPath())
		return nil
	})
	clientApp.AddCommand(send)
	clientApp.AddCommand(receiveHistory)
	clientApp.AddCommand(sendHistory)
	clientApp.AddCommand(reconnection)
}

// Register register a children grumble.App to root grumble.App
func Register(root *grumble.App) {
	root.AddCommand(clientCmd)
	clientApp.OnClose(func() error {
		if err := root.Close(); err != nil {
			return err
		}
		return nil
	})
}
