# Go-Assembly简明教程

## WebAssembly 简介

>   WebAssembly是一种新的编码方式，可以在现代的网络浏览器中运行；
>
>   它是一种低级的类汇编语言，具有紧凑的二进制格式，可以接近原生的性能运行，并为诸如C / C  ++等语言提供一个编译目标，以便它们可以在Web上运行；
>
>   它也被设计为可以与JavaScript共存，允许两者一起工作。  —— [MDN web docs - mozilla.org](https://developer.mozilla.org/zh-CN/docs/WebAssembly)

从 MDN 的介绍中，我们可以得出几个结论：

-   1）WebAssembly 是一种二进制编码格式，而不是一门新的语言；
-   2) WebAssembly 不是为了取代 JavaScript，而是一种补充（至少现阶段是这样），结合 WebAssembly 的性能优势，很大可能集中在对性能要求高（例如游戏，AI），或是对交互体验要求高（例如移动端）的场景；
-   3）C/C++ 等语言可以编译 WebAssembly 的目标文件，也就是说，其他语言可以通过编译器支持，而写出能够在浏览器前端运行的代码；

Go 语言在 1.11 版本(2018年8月) 加入了对 WebAssembly (Wasm) 的原生支持，使用 Go 语言开发  WebAssembly 相关的应用变得更加地简单，Go 语言的内建支持是 Go 语言进军前端的一个重要的里程碑。

在这之前，如果想使用 Go  语言开发前端，需要使用 [GopherJS](https://github.com/gopherjs/gopherjs)，GopherJS 是一个编译器，可以将 Go 语言转换成可以在浏览器中运行的 JavaScript 代码。**新版本的 Go 则直接将 Go 代码编译为 wasm  二进制文件，而不再需要转为 JavaScript 代码。**更巧的是，实现 GopherJS 和在 Go 语言中内建支持 WebAssembly  的是同一拨人。

**Go 语言实现的函数可以直接导出供 JavaScript 代码调用**

**同时，Go 语言内置了 [syscall/js](https://github.com/golang/go/tree/master/src/syscall/js) 包，可以在 Go 语言中直接调用 JavaScript 函数，包括对 DOM 树的操作!**

## Goland中使用Web-Assembly

在Goland中进行开发时需要进行一些简单的配置；

Goland官方文档也说明了如何配置：

https://www.jetbrains.com/help/go/webassembly-project.html

## Hello World

使用wasm实现的一个通知alert的例子；

见：https://github.com/JasonkayZK/go-assembly/tree/hello-world





