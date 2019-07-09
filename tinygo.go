package main

import (
	"fmt"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"os/exec"
)

// CompileToWASM - compile go code to wasm with tinygo
func CompileToWASM(c *cli.Context) {
	cmd := exec.Command("tinygo", "build", "-o", "build/out.wasm", "./src")

	fmt.Println(chalk.Magenta.Color("\nCompiling to WebAssembly...⌛"))

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	fmt.Println(
		chalk.Green.Color("\nCompiled successfully! ✨\n"),
		chalk.Blue.Color("\nNow start a dev server (gwa dev)\n"),
	)
}
