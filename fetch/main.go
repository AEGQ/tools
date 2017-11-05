package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.HideHelp = true
	app.HideVersion = true
	app.Commands = getCommands()
	app.Run(os.Args)
}

func getCommands() []cli.Command {
	cmds := []cli.Command{
		{
			Name:      "fetch",
			Usage:     "Fetch an image from any other docker host.",
			Flags:     getFetchFlags(),
			ArgsUsage: "user[:password]@IP:image[:tag]",
			Action: func(c *cli.Context) {
				fetchCmd(c)
			},
		},
	}
	return cmds
}
