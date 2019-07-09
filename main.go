package main

import (
	"github.com/urfave/cli"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "create-go-web-app"
	app.Usage = "Create a web app for writing frontend Go code"
	app.Author = "v1rtl (twitter.com/v1rtl)"
	app.Version = "0.0.1"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initialize a web app",
			Action: InitWebApp,
		},
		{
			Name:   "build",
			Usage:  "Compile Go to WebAssembly",
			Action: CompileToWASM,
		},
		{
			Name:   "dev",
			Usage:  "Run a development server",
			Action: RunDevServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port",
					Value: "8080",
					Usage: "Dev Server port",
				},
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
