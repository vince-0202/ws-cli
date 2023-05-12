package cmd

import (
	"github.com/desertbit/grumble"
	"github.com/vince-0202/ws-cli/cmd/client"
	"github.com/vince-0202/ws-cli/cmd/server"
)

// Version ws-cli version
const Version = "v0.0.0"

// Root grumble.App for ws-cli
var Root = grumble.New(&grumble.Config{
	Name:        "websocket",
	Description: "websocket testing tools",
	Flags: func(f *grumble.Flags) {
	},
})

func init() {
	Root.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("                        ___ ")
		a.Println(" _      _______   _____/ (_)")
		a.Println("| | /| / / ___/  / ___/ / / ")
		a.Println("| |/ |/ (__  )  / /__/ / /  ")
		a.Println("|__/|__/____/   \\___/_/_/  ")
		a.Println(" :: vince wang ::          (", Version, ")")
	})
	client.Register(Root)
	server.Register(Root)
	Root.AddCommand(&grumble.Command{
		Name: "version",
		Help: "print version",
		Run: func(c *grumble.Context) error {
			c.App.Println(Version)
			return nil
		},
	})
}
