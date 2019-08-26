package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CompileToWASM - compile go code to wasm with tinygo
func CompileToWASM() {

	cmd := exec.Command("tinygo", "build", "-o", "build/out.wasm", "./src")

	fmt.Println(chalk.Magenta.Color("\nCompiling to WebAssembly...⌛"))

	err := cmd.Run()

	if err != nil {
		log.Fatal("Failed compile to WebAssembly")
	}

	file, err := os.Open("build/out.wasm")

	if err != nil {
		log.Fatal("Failed to open WebAssembly output (.wasm)")
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		log.Fatal("Failed to load stats for WebAssembly output")
	}

	fmt.Print(
		chalk.Green.Color("\nCompiled successfully! ✨\n"),
		chalk.Yellow.Color("\nBundle Size: "),
		stat.Size()/1024,
		"Kb\n\n",
	)
}

// CompileToWASMCLI - the same as CompileToWASM but for CLI
func CompileToWASMCLI(c *cli.Context) {
	CompileToWASM()
}
