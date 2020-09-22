package main

import (
	"strconv"
	"syscall/js"
)

var (
	document = js.Global().Get("document")
	numEle   = document.Call("getElementById", "num")
	ansEle   = document.Call("getElementById", "ans")
	btnEle   = js.Global().Get("btn")
)

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	if num, err := strconv.Atoi(numEle.Get("value").String()); err == nil {
		ansEle.Set("innerHTML", js.ValueOf(fib(num)))
	}
	return nil
}

func main() {
	done := make(chan int, 0)
	btnEle.Call("addEventListener", "click", js.FuncOf(fibFunc))
	<-done
}
