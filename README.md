# Go-Assembly简明教程

## 回调函数(Callback Functions)

在 JavaScript 中，异步+回调是非常常见的，比如请求一个 Restful API，注册一个回调函数，待数据获取到，再执行回调函数的逻辑，这个期间程序可以继续做其他的事情，Go 语言可以通过协程实现异步。

假设 fib 的计算非常耗时，那么可以启动注册一个回调函数，待 fib 计算完成后，再把计算结果显示出来。

我们先修改 main.go，使得 fibFunc 支持传入回调函数。

```go
package main

import (
	"syscall/js"
	"time"
)

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1]
	go func() {
		time.Sleep(3 * time.Second)
		v := fib(args[0].Int())
		callback.Invoke(v)
	}()

	js.Global().Get("ans").Set("innerHTML", "Waiting 3s...")
	return nil
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<-done
}
```

-   假设调用 fibFunc 时，回调函数作为最后一个参数，那么**通过 args[len(args)-1] 便可以获取到该函数**。这与其他类型参数的传递并无区别。
-   使用 `go func()` 启动子协程，调用 fib 计算结果，计算结束后，调用回调函数 `callback`，并将计算结果传递给回调函数，使用 time.Sleep() 模拟 3s 的耗时操作。
-   计算结果出来前，先在界面上显示 `Waiting 3s...`

接下来我们修改 index.html，为按钮添加点击事件，调用 fibFunc：

```html
<html>
...
<body>
	<input id="num" type="number" />
	<button id="btn" onclick="fibFunc(num.value * 1, (v)=> ans.innerHTML=v)">Click</button>
	<p id="ans"></p>
</body>
</html>
```

-   为 btn 注册了点击事件，第一个参数是待计算的数字，从 num 输入框获取。
-   第二个参数是一个回调函数，将参数 v 显示在 ans 文本框中。

接下来，重新编译 main.go，访问 localhost:9999：

随便输入一个数字，点击 Click。页面会先显示 `Waiting 3s...`，3s过后显示计算结果；

