package main

import (
	"fmt"
	"io/ioutil"

	"rogchap.com/v8go"
)

func RunV8() {
	ctx, _ := v8go.NewContext()

	b, _ := ioutil.ReadFile("identity/runtime.js")
	runtime := string(b)
	b, _ = ioutil.ReadFile("identity/polyfills.js")
	polyfills := string(b)
	b, _ = ioutil.ReadFile("identity/vendor.js")
	vendor := string(b)
	b, _ = ioutil.ReadFile("identity/main.js")
	main := string(b)

	ctx.RunScript(runtime, "runtime.js")
	ctx.RunScript(polyfills, "polyfills.js")
	ctx.RunScript(vendor, "vendor.js")
	ctx.RunScript(main, "main.js")

	ctx.RunScript("const result = 18", "more.js")

	val, _ := ctx.RunScript("result", "value.js")
	fmt.Printf("addition result: %s", val)
}
