package client

import "github.com/desertbit/grumble"

var reconnection = &grumble.Command{
	Name: "reconnection",
	Help: "reconnection",
	Args: func(a *grumble.Args) {
	},
	Run: func(c *grumble.Context) error {
		if err := wsClient.Connection(); err != nil {
			return err
		}
		return nil
	},
}
