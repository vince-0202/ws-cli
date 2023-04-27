package server

import "github.com/desertbit/grumble"

var clientList = &grumble.Command{
	Name:    "clientList",
	Aliases: []string{"cl"},
	Help:    "list of clients",
	Run: func(c *grumble.Context) error {
		for _, addr := range wsServer.Clients() {
			if _, err := c.App.Println(addr); err != nil {
				return err
			}
		}
		return nil
	},
}
