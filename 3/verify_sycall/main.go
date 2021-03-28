//Wasming
// compile: GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {

	// Init Canvas stuff
	doc := js.Global().Get("document")

	done := make(chan struct{}, 0)

	var renderFrame js.Func
	var tmark float64
	var markCount = 0
	var tdiffSum float64
	var avg float64
	var oldIn = 0

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		now := args[0].Float()
		tdiffSum += now - tmark
		markCount++
		numIn, _ := strconv.Atoi((doc.Call("getElementById", "numberInput").Get("value").String()))
		if markCount > 5 {
			timems := tdiffSum / float64(markCount)
			if avg == 0 || oldIn != numIn {
				avg = timems
			} else {
				avg = (avg*9 + timems) / 10
			}
			doc.Call("getElementById", "output").Set("innerHTML", fmt.Sprintf("FPS: %.01f, %.01f ms, AVG %.01f ms", 1000/timems, timems, avg))
			tdiffSum, markCount = 0, 0
		}
		tmark = now
		oldIn = numIn
		res := doWork(doc, numIn)
		doc.Call("getElementById", "result").Set("innerHTML", res)
		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done

}

func doWork(doc js.Value, numIn int) int {
	var res = 0
	for i := 0; i < numIn; i++ {
		numIn, _ := strconv.Atoi((doc.Call("getElementById", "numberInput").Get("value").String()))
		res += numIn
	}
	return res
}
