# Go-Assembly简明教程

## 使用Web Assembly实现Hello-World

### 使用Goland前配置

在Goland中进行开发时需要进行一些简单的配置；

Goland官方文档也说明了如何配置：

https://www.jetbrains.com/help/go/webassembly-project.html

### 实现步骤

 第一步，新建文件 main.go，使用 js.Global().get(‘alert’) 获取全局的 alert 对象，通过 Invoke 方法调用；

等价于在 js 中调用 `window.alert("Hello World")`。 

```go
// main.go
package main

import "syscall/js"

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")
}
```

第二步，将 main.go 编译为 static/main.wasm

>   如果启用了 `GO MODULES`，则需要使用 go mod init 初始化模块，或设置 GO111MODULE=auto。

```bash
$ GOOS=js GOARCH=wasm go build -o static/main.wasm
```

第三步，拷贝 wasm_exec.js (JavaScript 支持文件，加载 wasm 文件时需要) 到 static 文件夹

```bash
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static
```

第四步，创建 index.html，引用 `static/main.wasm` 和 `static/wasm_exec.js`。

```html
<html>
<script src="static/wasm_exec.js"></script>
<script>
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("static/main.wasm"), go.importObject)
		.then((result) => go.run(result.instance));
</script>
</html>
```

第五步，使用 goexec 启动 Web 服务

>   如果没有安装 goexec，可用 `go get -u github.com/shurcooL/goexec` 安装，需要将 $GOBIN 或 $GOPATH/bin 加入环境变量

当前的目录结构如下：

```bash
demo/
   |--static/
      |--wasm_exec.js
      |--main.wasm
   |--main.go
   |--index.html
```

```bash
goexec "http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))"
```

浏览器访问 localhost:9999，则会有一个弹出窗口，上面写着 *Hello World!*；

为了避免每次编译都需要输入繁琐的命令，可将这个过程写在 `Makefile` 中：

```makefile
all: static/main.wasm static/wasm_exec.js
	goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'

static/wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" static

static/main.wasm : main.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o static/main.wasm .
```

 这样一个敲一下 make 就够了！

