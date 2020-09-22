# Go-Assembly简明教程

## 注册函数(Register Functions)

在 Go 语言中调用 JavaScript 函数是一方面，另一方面，如果仅仅是使用 WebAssembly 替代性能要求高的模块，那么就需要注册函数，以便其他 JavaScript 代码调用。

假设我们需要注册一个计算斐波那契数列的函数，可以这么实现：

```go
// main.go
package main

import "syscall/js"

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(fib(args[0].Int()))
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<-done
}
```

-   fib 是一个普通的 Go 函数，通过递归计算第 i 个斐波那契数，接收一个 int 入参，返回值也是 int。
-   定义了 fibFunc 函数，为 fib 函数套了一个壳，从 args[0] 获取入参，计算结果用 js.ValueOf 包装，并返回。
-   使用 js.Global().Set() 方法，将注册函数 fibFunc 到全局，以便在浏览器中能够调用。

`js.Value` 可以将 Js 的值转换为 Go 的值，比如 args[0].Int()，则是转换为 Go 语言中的整型。

`js.ValueOf`，则用来将 Go 的值，转换为 Js 的值。

另外，注册函数的时候，使用 js.FuncOf 将函数转换为 `Func` 类型，只有 Func 类型的函数，才能在 JavaScript 中调用。可以认为这是 Go 与 JavaScript 之间的接口/约定。

`js.Func()` 接受一个函数类型作为其参数，该函数的定义必须是：

```go
func(this Value, args []Value) interface{}
// this 即 JavaScript 中的 this
// args 是在 JavaScript 中调用该函数的参数列表。
// 返回值需用 js.ValueOf 映射成 JavaScript 的值
```

在 main 函数中，创建了信道(chan) done，阻塞主协程(goroutine)；

fibFunc 如果在 JavaScript 中被调用，会开启一个新的子协程执行：

>   A wrapped function triggered during a call from Go to JavaScript gets executed on the same goroutine.
>
>   A wrapped function triggered by  JavaScript’s event loop gets executed on an extra goroutine.  —— [FuncOf - golang.org](https://golang.org/pkg/syscall/js/#FuncOf)

接下来，修改之前的 index.html，在其中添加一个输入框(num)，一个按钮(btn) 和一个文本框(ans，用来显示计算结果)，并给按钮添加了一个点击事件，调用 fibFunc，并将计算结果显示在文本框(ans)中。

```html
<html>
...
<body>
	<input id="num" type="number" />
	<button id="btn" onclick="ans.innerHTML=fibFunc(num.value * 1)">Click</button>
	<p id="ans">1</p>
</body>
</html>
```

使用之前的命令重新编译 main.go，并在 9999 端口启动 Web 服务，如果我们已经将命令写在 Makefile 中了，只需要运行 `make` 即可；

接下来访问 localhost:9999，可以看到效果：

输入一个数字，点击`Click`，计算结果显示在输入框下方；