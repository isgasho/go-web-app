package main

import (
	"fmt"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"net/http"
)

// RunDevServer - launches a server for hosting html page with wasm_exec and fetched wasm binaries
func RunDevServer(c *cli.Context) {
	port := c.String("port")

	fmt.Println(
		chalk.Bold.TextStyle("\nListening on http://localhost:"+port),
		chalk.Yellow.Color("\n\nShutdown server: Ctrl + C"),
	)

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":"+port, nil)
}
