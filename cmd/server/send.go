package server

import "github.com/desertbit/grumble"

var send = &grumble.Command{
	Name:    "send",
	Aliases: []string{"sm", "s"},
	Help:    "send a message to the websocket server",
	Args: func(a *grumble.Args) {
		a.String("msg", "msg")
	},
	Flags: func(f *grumble.Flags) {
		f.String("c", "client", "", "client")
	},
	Run: func(c *grumble.Context) error {
		if c.Flags.String("client") == "" {
			if err := wsServer.SendTextToAllConnection(c.Args.String("msg")); err != nil {
				return err
			}
		} else {
			if err := wsServer.SendText(c.Flags.String("client"), c.Args.String("msg")); err != nil {
				return err
			}
		}
		return nil
	},
}
