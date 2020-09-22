# Go-Assembly简明教程

## 操作DOM

在[注册函数(Register Functions)](https://github.com/JasonkayZK/go-assembly/tree/register-function)例子中，仅仅是注册了全局函数 fibFunc，事件注册，调用，对 DOM 元素的操作都是在 HTML中通过原生的 JavaScript 函数实现的。

这些事情，能不能全部在 Go 语言中完成呢？答案可以。

首先修改 index.html，删除事件注册部分和 对 DOM 元素的操作部分。

```html
<html>
...
<body>
	<input id="num" type="number" />
	<button id="btn">Click</button>
	<p id="ans">1</p>
</body>
</html>
```

修改 main.go：

```go
package main

import (
	"strconv"
	"syscall/js"
)

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

var (
	document = js.Global().Get("document")
	numEle   = document.Call("getElementById", "num")
	ansEle   = document.Call("getElementById", "ans")
	btnEle   = js.Global().Get("btn")
)

func fibFunc(this js.Value, args []js.Value) interface{} {
	v := numEle.Get("value")
	if num, err := strconv.Atoi(v.String()); err == nil {
		ansEle.Set("innerHTML", js.ValueOf(fib(num)))
	}
	return nil
}

func main() {
	done := make(chan int, 0)
	btnEle.Call("addEventListener", "click", js.FuncOf(fibFunc))
	<-done
}
```

-   通过 `js.Global().Get("btn")` 或 `document.Call("getElementById", "num")` 两种方式获取到 DOM 元素。
-   btnEle 调用 `addEventListener` 为 btn 绑定点击事件 fibFunc。
-   在 fibFunc 中使用 `numEle.Get("value")` 获取到 numEle 的值（字符串），转为整型并调用 fib 计算出结果。
-   ansEle 调用 `Set("innerHTML", ...)` 渲染计算结果。

重新编译 main.go，访问 localhost:9999，效果与之前是一致的！
