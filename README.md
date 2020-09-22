# Go-Assembly简明教程

## WebAssembly 简介

>   WebAssembly是一种新的编码方式，可以在现代的网络浏览器中运行；
>
>   它是一种低级的类汇编语言，具有紧凑的二进制格式，可以接近原生的性能运行，并为诸如C / C  ++等语言提供一个编译目标，以便它们可以在Web上运行；
>
>   它也被设计为可以与JavaScript共存，允许两者一起工作。  —— [MDN web docs - mozilla.org](https://developer.mozilla.org/zh-CN/docs/WebAssembly)

从 MDN 的介绍中，我们可以得出几个结论：

-   1）WebAssembly 是一种二进制编码格式，而不是一门新的语言；
-   2）WebAssembly 不是为了取代 JavaScript，而是一种补充（至少现阶段是这样），结合 WebAssembly 的性能优势，很大可能集中在对性能要求高（例如游戏，AI），或是对交互体验要求高（例如移动端）的场景；
-   3）C/C++ 等语言可以编译 WebAssembly 的目标文件，也就是说，其他语言可以通过编译器支持，而写出能够在浏览器前端运行的代码；

Go 语言在 1.11 版本(2018年8月) 加入了对 WebAssembly (Wasm) 的原生支持，使用 Go 语言开发  WebAssembly 相关的应用变得更加地简单，Go 语言的内建支持是 Go 语言进军前端的一个重要的里程碑。

在这之前，如果想使用 Go  语言开发前端，需要使用 [GopherJS](https://github.com/gopherjs/gopherjs)，GopherJS 是一个编译器，可以将 Go 语言转换成可以在浏览器中运行的 JavaScript 代码。

**新版本的 Go 则直接将 Go 代码编译为 wasm  二进制文件，而不再需要转为 JavaScript 代码。**

更巧的是，实现 GopherJS 和在 Go 语言中内建支持 WebAssembly  的是同一拨人。

**Go 语言实现的函数可以直接导出供 JavaScript 代码调用**

**同时，Go 语言内置了 [syscall/js](https://github.com/golang/go/tree/master/src/syscall/js) 包，可以在 Go 语言中直接调用 JavaScript 函数，包括对 DOM 树的操作!**

<BR/>

## Goland中使用Web-Assembly

在Goland中进行开发时需要进行一些简单的配置；

Goland官方文档也说明了如何配置：

https://www.jetbrains.com/help/go/webassembly-project.html

<BR/>

## Hello World

使用wasm实现的一个通知alert的例子；

见：https://github.com/JasonkayZK/go-assembly/tree/hello-world

<BR/>

## 注册函数(Register Functions)

在Go中注册JS函数，以便其他 JavaScript 代码调用 ；

见：https://github.com/JasonkayZK/go-assembly/tree/register-function

<BR/>

## 操作DOM

在Go中操作DOM；

见：https://github.com/JasonkayZK/go-assembly/tree/operate-dom

<BR/>

## 回调函数(Callback Functions)

 在 JavaScript 中，异步+回调是非常常见的，比如请求一个 Restful API，注册一个回调函数，待数据获取到，再执行回调函数的逻辑，这个期间程序可以继续做其他的事情；

Go 语言可以通过协程实现异步，进而简化编程；

见：https://github.com/JasonkayZK/go-assembly/tree/callback-function

<BR/>

## 进一步的尝试

### 6.1 工具框架

-   WebAssembly 的二进制分析工具 [WebAssembly Code Explorer](https://wasdk.github.io/wasmcodeexplorer/)
-   使用NodeJs 或浏览器测试 Go Wasm 代码 [Github Wiki](https://github.com/golang/go/wiki/WebAssembly#executing-webassembly-with-nodejs)
-   借鉴 Vue 实现的 Golang WebAssembly 前端框架 [Vugu](https://www.vugu.org/doc/start)，完全使用 Go，不用写任何的 JavaScript 代码。

### 6.2 Demo/项目

-   使用 Go Assembly 前端渲染的一些[例子](https://stdiopt.github.io/gowasm-experiments/)
-   [jsgo](https://github.com/dave/jsgo) 这个项目汇聚一些小而精的项目，包括 [2048](https://jsgo.io/hajimehoshi/ebiten/examples/2048)，[俄罗斯方块](https://jsgo.io/hajimehoshi/ebiten/examples/blocks)等游戏，还有证明 Go 可以完整开发前端项目的 [TodoMVC](https://jsgo.io/dave/todomvc)

### 6.3 相关文档

-   [syscall/js 官方文档 - golang.org](https://golang.org/pkg/syscall/js)
-   [Go WebAssembly 官方文档 - github.com](https://github.com/golang/go/wiki/WebAssembly)

<BR/>

## 附录

本文转自：

-   [Go WebAssembly (Wasm) 简明教程](https://geektutu.com/post/quick-go-wasm.html)

