# go-web-app

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/talentlessguy/create-go-web-app.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/talentlessguy/create-go-web-app?style=flat-square)](https://goreportcard.com/report/github.com/talentlessguy/create-go-web-app)

Simple CLI for setting up Go WebAssembly frontend app.

## What's included 📦

* Dev Server
* [TinyGo](https://tinygo.org) for small WebAssembly output
* Git setup
* Glue files (`index.html` + `wasm_exec.js`)

## Requirements ✅

* Go 1.12+ (I have Go 1.12.4)
* TinyGo installed. Install [here](https://tinygo.org/getting-started)
* Browser that supports WebAssembly (I use Mozilla Dev Edition 68)

## Install 🔄

### Using Bash script

```sh
curl https://raw.githubusercontent.com/talentlessguy/go-web-app/master/install.sh | bash
```

This will install `gwa` into `/usr/local/bin` so be sure that `/usr/local/bin` is in your `$PATH`.

### With `go get`

```sh
go get github.com/talentlessguy/go-web-app
```

Then use as `go-web-app` command instead of `gwa`

## Commands 💻

### `gwa init`

Initialize a project.

#### Project tree

`out.wasm` is generated when building. Other files are automatically added.

```text
├── build
│   └── out.wasm
├── src
│   └── main.go
├── wasm_exec.js
└── index.html
```

### `gwa dev --port <port>`

Launches a development server with specified port.

Default port is **8080**.

After launching a server, you should go to `http://localhost:<port>`

### `gwa build`

Compiles go code to WebAssembly. Compiled `out.wasm` file could be found in `build` folder.

Everything in `src` compiles to `build`, every go file.
