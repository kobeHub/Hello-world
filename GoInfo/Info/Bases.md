# Golang Bases

[TOC]

## 1. Why Golang

golang 是一个开源，编译的，静态类型的编程语言．其主要目的是更加简单的开发高可用性，可扩展的web应用．虽然有很多可以完成相同工作的语言，但是与之相比，golang有以下优势:

+ 并发性是该语言固有的部分,更加简易的进行多进程编写,通过`Goroutines` `channels` 实现
+ Golang是编译性的语言,可以产生二进制文件,对于js,python等动态语言不具备该特征
+ 支持静态链接,所有的go代码可以静态连接到一个巨大的二进制库中,可以快速部署到云服务器而不需担心依赖问题 

## ２. Workspace and directory structure

 首先需要建立一个go的工作目录,可以直接在当前终端`export GOPATH=/path/` 作为工作目录,对于golang编译产生的可执行文件,位于当前工作目录下的`bin`文件夹,可以在path中引入该环境变量

`export PATH=${PATH}:${GOPATH}/bin` 

go项目的基本目录结构如下:

```shell
├── bin
│   └── hello
└── src
    └── github.com
        └── nf
            ├── hello
            │   └── echo.go
            └── string
                ├── reverse.go
                └── reverse_test.go
```

+ `src` :

  源代码存放的目录,可以包含多个子文件夹,每个文件夹包含多个包,对于`hello` `string` 是直接可以被引用的包名

+ `bin`:

  可执行文件的位置

对于程序的入口,需要以`package main` 作为包名,进行`go build` `go install` 后生成相应的二进制文件,也可以直接使用`go run xxxx.go` 以类似脚本的形式进行运行,go run 命令只能接受一个命令源码文件以及若干个库源码文件（必须同属于 main 包）作为文件参数，且**不能接受测试源码文件**。它在执行时会检查源码文件的类型。如果参数中有多个或者没有命令源码文件，那么 go run 命令就只会打印错误提示信息并退出，而不会继续执行.

## 3. Syntax

### 1. 代码结构

golang 为静态类型语言,对于每一个变量必须赋予其相应类型,golang相邻的语句间不需要添加`;` 作为间隔,使用{}表示语句块

```go
package main //1

import "fmt" //2

func main() { //3  
    fmt.Println("Hello World") //4
}
```

+ `package main` 每个源代码文件必须以一个包名为开始,标识其所属的包,`main` 包名是程序入口包名

+ `import` 引入所需要的包名,多个包名可以采用()引入,并且可以更名

  ```go
  import (
  	fmt "fmt"
      Str "string"
      "time"
      "math/rand"
  )  // 可以在每个包名后添加;如果不加编译器会自动添加,也可以写为一行,添加分号
  ```

+ `func main` 程序入口

###2. 变量

可以使用多种方式声名变量,可以使用`var`关键字,采用`var variable_name type` 的语法进行变量声明,也可以在声明变量时赋予初始值;在赋予初始值的同时可以省去变量的类型,根据变量值来推测其类型;还可以使用更为优雅的方式,同时声明多个变量;在代码块内部,可以使用短声明的方式声明一个变量`a := 12.0` ,但是不可以用在函数外部

```go
package main

import (
	"fmt"
    "math"
)

var (
	var1 int
    var2 string
    name = "kobe"
    age = 39
)

func main() {
    var test string = "It's just a test"
    var2 = test
    var1 = age
    a, b := 12.6, 22.0
    a, b = b, a
    
    fmt.Println(var1, var2, test, name, age, math.Sqrt(a))
}
```

**注意:使用()进行多个变量声明或者包引入时,每个item间可以用`;`隔开也可以不用,不可以用`,` ** 

go为静态语言不可以将不同类型的数据赋值给不同类型的变量

### 3. 基本类型

+ bool

+ Numeric Types
  + int8, int16, int32, int64, int
  + uint8, uint16, uint32, uint64, uint
  + float32, float64

    complex64, complex128
  + byte
  + rune

+ string

#### Signed integers

| Type  | size                | range                   |
| ----- | ------------------- | ----------------------- |
| int8  | 8 bits              | -128-127                |
| int16 | 16 bits             | -32768-32767            |
| int32 | 32 bits             | -2147483648- 2147483647 |
| int64 | 64 bits             |                         |
| int   | depends on platform | 64 bits                 |

可以使用`fmt.Printf` 中的格式化输出`%T`获取一个变量或者数据的类型，也可以使用`unsafe`包中的`Sizeof（）` 函数获得变量的大小(Bytes)

```go
package main

import (
  "fmt"
  "unsafe"
)

func main() {
  a, b := 'a', 12
  c, d, e := 123, "Test", "go"

  fmt.Println(a, b, c, d, e)
  fmt.Println("Type test:")
  fmt.Printf("a: %T, size: %d, b: %T, size: %d", a, unsafe.Sizeof(a), b, unsafe.Sizeof(b))

}


97 12 123 Test go
Type test:
a: int32, size: 4, b: int, size: 8%  
```

#### Unsigned int

与有符号数类似

#### Floating point types

浮点数，包含32位以及64位的类型

#### Complex types

+ complex64: 实部和虚部都是32为浮点数
+ complex128：实部和虚部都是64位浮点数

可以使用内置函数初始化一个复数,传入的参数类型需要相同，得到的复数的类型根据传入的参数决定

```go
func complex(r, i FloatType) ComplexType
```

或者可以直接使用变量声明的方式：

```go
c := 1 + 2i
```

其他两种数据类型byte 对应于uint8, rune对应于int32,字符串类型必须使用`“”` 括起来，使用单引号对应于byte

### 4. 类型转换

go语言不支持不同类型的number的直接运算，需要进行显式的类型转换

```go
i := 2
j := 2.2
sum := i + int(j)   // 4

var j float64 = float64(12)  // 必须加后面的强制类型转换
```

### 5.常量使用

使用const关键字定义常量，注意常量的处理是在编译时进行的，不可以在运行时给常量赋值，类似于C中的`#define` 。所以不可以给一个常量赋值为一个函数的返回值

在go语言中任意包含在双括号中的string，都是字符串常量，该常量没有类型`untyped` 

```go
const hello = "The string is untyped"
```

这似乎和之前的变量声明有所矛盾，如果"string" 没有类型，**对于`var a = "cds"`那又是如何根据常量“cds”向a进行类型传递呢？**

这是因为无类型的常量有一个默认的类型与之关联，只有代码真正需要使用该默认类型时才会使用，

对于以上声明操作，由于变量a需要一个类型所以使用了默认类型`string` 

```go
const typed string = "cdsvds"   // 声明有类型的常量

func main() {  
        var defaultName = "Sam" //allowed
        type myString string
        var customName myString = "Sam" //allowed
        customName = defaultName //not allowed

}
```

**Golang是强类型语言，所以以上代码会报错**。虽然`myString` 是`string` 的一个别名，但是强类型不允许不同类型的变量相互赋值，所以该操作不被允许

与之相似bool常量操作与string常量相同，别名类型变量赋值仍会报错。

**数字常量**

数字常量包含所有int, float, complex.对于数字常量而言，也是没有类型的，所以同一个常量可以赋给各种类型的变量，并且根据需求进行相应的类型转换。对于单一的`const a = 5` 而言其默认类型是int, float64, complex128(根据平台判断)