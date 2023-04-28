package cmd

import (
	"github.com/desertbit/grumble"
	"github.com/vince-0202/ws-cli/cmd/client"
	"github.com/vince-0202/ws-cli/cmd/server"
)

// Root grumble.App for ws-cli
var Root = grumble.New(&grumble.Config{
	Name:        "websocket",
	Description: "websocket testing tools",
	Flags: func(f *grumble.Flags) {
	},
})

func init() {
	client.Register(Root)
	server.Register(Root)
}
