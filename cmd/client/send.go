package client

import (
	"github.com/desertbit/grumble"
)

var send = &grumble.Command{
	Name:    "send",
	Aliases: []string{"s"},
	Help:    "send a message to the websocket server",
	Args: func(a *grumble.Args) {
		a.String("msg", "msg")
	},
	Run: func(c *grumble.Context) error {

		if err := wsClient.SendText(c.Args.String("msg")); err != nil {
			return err
		}
		return nil
	},
}
