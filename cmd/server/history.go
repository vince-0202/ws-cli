package server

import "github.com/desertbit/grumble"

var sendHistory = &grumble.Command{
	Name: "sendHistory",
	Help: "history of all messages sent",
	Run: func(c *grumble.Context) error {
		for _, history := range wsServer.HistoryOfSend() {
			c.App.Println(history.ToString())
		}
		return nil
	},
}

var receiveHistory = &grumble.Command{
	Name: "received",
	Help: "history of all messages receive",
	Run: func(c *grumble.Context) error {
		for _, history := range wsServer.HistoryOfReceive() {
			c.App.Println(history.ToString())
		}
		return nil
	},
}
