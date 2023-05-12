package client

import "github.com/desertbit/grumble"

var reconnection = &grumble.Command{
	Name:    "reconnection",
	Aliases: []string{"rc"},
	Help:    "close and reconnection to websocket server",
	Args: func(a *grumble.Args) {
	},
	Run: func(c *grumble.Context) error {
		_ = wsClient.CloseConnection()
		if err := wsClient.Connection(); err != nil {
			return err
		}
		return nil
	},
}
