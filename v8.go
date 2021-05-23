package main

import (
	"fmt"

	"rogchap.com/v8go"
)

func RunV8() {
	ctx, _ := v8go.NewContext()
	ctx.RunScript("const add = (a, b) => a + b", "math.js")
	ctx.RunScript("const result = add(13, 4)", "main.js")
	val, _ := ctx.RunScript("result", "value.js")
	fmt.Printf("addition result: %s", val)
}
