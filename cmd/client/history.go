package client

import "github.com/desertbit/grumble"

var sendHistory = &grumble.Command{
	Name: "sendHistory",
	Help: "history of all messages sent",
	Args: func(a *grumble.Args) {
	},
	Run: func(c *grumble.Context) error {
		for _, history := range wsClient.HistoryOfSend() {
			c.App.Println(history.ToString())
		}
		return nil
	},
}

var receiveHistory = &grumble.Command{
	Name: "received",
	Help: "history of all messages receive",
	Args: func(a *grumble.Args) {
	},
	Run: func(c *grumble.Context) error {
		for _, history := range wsClient.HistoryOfReceive() {
			c.App.Println(history.ToString())
		}
		return nil
	},
}
