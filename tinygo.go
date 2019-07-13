package main

import (
	"fmt"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"os"
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

	file, err := os.Open("build/out.wasm")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	stat, err := file.Stat()

	fmt.Println(
		chalk.Green.Color("\nCompiled successfully! ✨\n"),
		chalk.Yellow.Color("\nBundle Size:"),
		stat.Size()/1024,
		"Kb\n",
		chalk.Blue.Color("\nNow start a dev server (gwa dev)\n"),
	)
}
