package main

import (
	"fmt"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

// Create a project directory
func createDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}
	os.Chdir(dir)
}

func initGitRepo() {
	cmd := exec.Command("git", "init")

	fmt.Println(chalk.Blue.Color("Initializing a Git repository..."))
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func setupGlue() {

	fmt.Println(chalk.Green.Color("Creating WebAssembly glue..."))

	// wasm_exec.js file

	wasmExec, err := os.Create("wasm_exec.js")

	if err != nil {
		panic(err)
	}

	defer wasmExec.Close()

	wasmExec.WriteString(`(() => { const e = "undefined" != typeof process; if (e) { global.require = require, global.fs = require("fs"); const e = require("crypto"); global.crypto = { getRandomValues(t) { e.randomFillSync(t) } }, global.performance = { now() { const [e, t] = process.hrtime(); return 1e3 * e + t / 1e6 } }; const t = require("util"); global.TextEncoder = t.TextEncoder, global.TextDecoder = t.TextDecoder } else { if ("undefined" != typeof window) window.global = window; else { if ("undefined" == typeof self) throw new Error("cannot export Go (neither window nor self is defined)"); self.global = self } let e = ""; global.fs = { constants: { O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1 }, writeSync(t, n) { const r = (e += s.decode(n)).lastIndexOf("\n"); return -1 != r && (console.log(e.substr(0, r)), e = e.substr(r + 1)), n.length }, write(e, t, s, n, r, i) { if (0 !== s || n !== t.length || null !== r) throw new Error("not implemented"); i(null, this.writeSync(e, t)) }, open(e, t, s, n) { const r = new Error("not implemented"); r.code = "ENOSYS", n(r) }, fsync(e, t) { t(null) } } } const t = new TextEncoder("utf-8"), s = new TextDecoder("utf-8"); var n = []; if (global.Go = class { constructor() { this._callbackTimeouts = new Map, this._nextCallbackTimeoutID = 1; const e = () => new DataView(this._inst.exports.memory.buffer), r = t => { const s = e().getFloat64(t, !0); if (0 === s) return; if (!isNaN(s)) return s; const n = e().getUint32(t, !0); return this._values[n] }, i = (t, s) => { if ("number" == typeof s) return isNaN(s) ? (e().setUint32(t + 4, 2146959360, !0), void e().setUint32(t, 0, !0)) : 0 === s ? (e().setUint32(t + 4, 2146959360, !0), void e().setUint32(t, 1, !0)) : void e().setFloat64(t, s, !0); switch (s) { case void 0: return void e().setFloat64(t, 0, !0); case null: return e().setUint32(t + 4, 2146959360, !0), void e().setUint32(t, 2, !0); case !0: return e().setUint32(t + 4, 2146959360, !0), void e().setUint32(t, 3, !0); case !1: return e().setUint32(t + 4, 2146959360, !0), void e().setUint32(t, 4, !0) }let n = this._refs.get(s); void 0 === n && (n = this._values.length, this._values.push(s), this._refs.set(s, n)); let r = 0; switch (typeof s) { case "string": r = 1; break; case "symbol": r = 2; break; case "function": r = 3 }e().setUint32(t + 4, 2146959360 | r, !0), e().setUint32(t, n, !0) }, o = (e, t, s) => new Uint8Array(this._inst.exports.memory.buffer, e, t), l = (e, t, s) => { const n = new Array(t); for (let s = 0; s < t; s++)n[s] = r(e + 8 * s); return n }, a = (e, t) => s.decode(new DataView(this._inst.exports.memory.buffer, e, t)), c = Date.now() - performance.now(); this.importObject = { env: { io_get_stdout: function () { return 1 }, resource_write: function (t, r, i) { if (1 == t) for (let t = 0; t < i; t++) { let i = e().getUint8(r + t); if (13 == i); else if (10 == i) { let e = s.decode(new Uint8Array(n)); n = [], console.log(e) } else n.push(i) } else console.error("invalid file descriptor:", t) }, "runtime.ticks": () => c + performance.now(), "runtime.sleepTicks": e => { setTimeout(this._inst.exports.go_scheduler, e) }, "syscall/js.stringVal": (e, t, s) => { const n = a(t, s); i(e, n) }, "syscall/js.valueGet": (e, t, s, n) => { let o = a(s, n), l = r(t), c = Reflect.get(l, o); i(e, c) }, "syscall/js.valueSet": (e, t, s, n) => { const i = r(e), o = a(t, s), l = r(n); Reflect.set(i, o, l) }, "syscall/js.valueIndex": (e, t, s) => { i(e, Reflect.get(r(t), s)) }, "syscall/js.valueCall": (t, s, n, o, c, u, f) => { const d = r(s), h = a(n, o), w = l(c, u); try { const s = Reflect.get(d, h); i(t, Reflect.apply(s, d, w)), e().setUint8(t + 8, 1) } catch (s) { i(t, s), e().setUint8(t + 8, 0) } }, "syscall/js.valueNew": (t, s, n, o, a) => { const c = r(s), u = l(n, o); try { i(t, Reflect.construct(c, u)), e().setUint8(t + 8, 1) } catch (s) { i(t, s), e().setUint8(t + 8, 0) } }, "syscall/js.valueLength": e => r(e).length, "syscall/js.valuePrepareString": (s, n) => { const o = String(r(n)), l = t.encode(o); i(s, l), ((t, s) => { e().setUint32(t + 0, s, !0), e().setUint32(t + 4, Math.floor(s / 4294967296), !0) })(s + 8, l.length) }, "syscall/js.valueLoadString": (e, t, s, n) => { const i = r(e); o(t, s).set(i) } } } } async run(e) { this._inst = e, this._values = [NaN, 0, null, !0, !1, global, this._inst.exports.memory, this], this._refs = new Map, this._callbackShutdown = !1, this.exited = !1; new DataView(this._inst.exports.memory.buffer); for (; ;) { const e = new Promise(e => { this._resolveCallbackPromise = (() => { if (this.exited) throw new Error("bad callback: Go program has already exited"); setTimeout(e, 0) }) }); if (this._inst.exports.cwa_main(), this.exited) break; await e } } _resume() { if (this.exited) throw new Error("Go program has already exited"); this._inst.exports.resume(), this.exited && this._resolveExitPromise() } _makeFuncWrapper(e) { const t = this; return function () { const s = { id: e, this: this, args: arguments }; return t._pendingEvent = s, t._resume(), s.result } } }, e) { 3 != process.argv.length && (process.stderr.write("usage: go_js_wasm_exec [wasm binary] [arguments]\n"), process.exit(1)); const e = new Go; WebAssembly.instantiate(fs.readFileSync(process.argv[2]), e.importObject).then(t => (process.on("exit", t => { 0 !== t || e.exited || (e._callbackShutdown = !0) }), e.run(t.instance))).catch(e => { throw e }) } })()`)

	// index.html file

	indexHTML, err := os.Create("index.html")

	if err != nil {
		panic(err)
	}

	defer indexHTML.Close()

	indexHTML.WriteString(`<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>Go WebAssembly App</title>
	</head>
	<body>
		<script src="wasm_exec.js"></script>
		<script>
			// Polyfill
			if (!WebAssembly.instantiateStreaming) {
				WebAssembly.instantiateStreaming = async (res, obj) => (
					await WebAssembly.instantiate(await (await res).arrayBuffer(), obj)
				)
			}
			const go = new Go()
			WebAssembly.instantiateStreaming(fetch('/build/out.wasm'), go.importObject).then((result) => {
				go.run(result.instance)
				WebAssembly.instantiate(result.module, go.importObject)
			})
		</script>
	</body>
	</html>`)
}

func helloWorld() {

	createDir("src")

	os.Chdir("src")

	// Go hello world file

	mainGo, err := os.Create("main.go")

	fmt.Println(chalk.Cyan.Color("Setting up Hello World..."))

	if err != nil {
		panic(err)
	}

	defer mainGo.Close()

	mainGo.WriteString(`package main

import "fmt"

func main() {
	fmt.Println("Hello WASM!")
}`)
}

func createOutputDir() {
	os.Chdir("..")
	createDir("build")
}

// InitWebApp - create a project with files
func InitWebApp(c *cli.Context) {

	appName := c.Args().First()

	// Create a project directory
	createDir(appName)
	os.Chdir(appName)
	initGitRepo()
	setupGlue()
	helloWorld()
	createOutputDir()

	fmt.Println(
		"Go Web app is ready!",
		"\nTo get started, go to project directory\n\n",
		chalk.Blue.Color("cd "+appName),
		"\n\nThen, compile to WebAssembly\n\n",
		chalk.Blue.Color("cgwa build"),
		"\n\nAnd start a development server\n\n",
		chalk.Blue.Color("cgwa dev"),
		chalk.Bold.TextStyle("\n\nHappy coding!\n"),
	)
}
