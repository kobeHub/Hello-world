# Object-oriented Programming

[TOC]

## 1. Go不是纯面向对象的语言

Golang不是一个纯面向对象的语言，摘自Golang [FAQ](This excerpt taken from Go's [FAQs](https://golang.org/doc/faq#Is_Go_an_object-oriented_language) answers the question of whether Go is Object Oriented.) ,解释了golang是否为面向对象的语言。

```
Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).  
```

也就是说使用类型以及menthods可以实现面向对象编程的类似功能，但是不存在类型的继承。使用了`interface`实现了不一样的解决方案，更加易用而且更加通用。

可以使用`struct`代替类,如果需要使用可见性约束的类，可以定义包内类型，然后定义`exported methods` ，定义`New() T`作为构造器的模拟，同时使用其他方法，使得接收器为该类型变量：

```go
// Define the struct with all fields exported
type employee struct {
  firstName, lastName string
  totalLeaves, leavesTaken int
}

// Exported method with a receiver of Employee
func (e employee) LeavesRemaining() {
  fmt.Printf("%s %s has %d leaves remaining.\n",
              e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

// New function as a Constructor
func New(firstName, lastName string, totalLeaves, leavesTaken int)  employee {
  e := employee {firstName, lastName, totalLeaves, leavesTaken}
  return e
}

// Property of employee
func (e employee) GetFirstName() string {
  return e.firstName
}

func (e employee) GetLastName() string {
  return e.lastName
}

func (e employee) GetTotal() int {
  return e.totalLeaves
}
```

## 2. 使用组合代替继承

Golang不支持继承机制，但是支持组合`composition`。组合的一般定义是能将多个部分放在一起。在Golang中使用接阔的嵌套实现组合机制，

### 2.1 结构体的嵌套

在一个结构体中使用另一个结构体作为其域可以实现组合的机制，同时借助匿名域可以直接使用外部类型调用内部类型的方法。使用匿名域时只需要指明类型即可，不需要给一个域命名，但是匿名域每一个类型只可以有一个。

```go
// author of the post
type author struct {
  firstName, lastName string
  bio string
}

// get the author full name
func (a author) AuthorName() string {
  return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// struct of post, `author` field promoted
// so Post type can use `AuthorName` directly
type Post struct {
  Title string
  Content string
  author
}

func (p Post) Detail() {
  fmt.Println("Post:", p.Title)
  fmt.Println("Content:", p.Content)
  fmt.Println("Author:", p.AuthorName())
  fmt.Println("Bio", p.bio)
}
```

### 2.2 嵌套的切片类型

可以在一个结构体内部定义另一个结构体的切片作为该结构体的域，但是注意如果需要访问该切片field，就不可以将其定义为匿名域

```go
type website struct {  
        []post
}
func (w website) contents() {  
    fmt.Println("Contents of Website\n")
    for _, v := range w.posts {
        v.details()
        fmt.Println()
    }
}
```

编译器会报错：

```shell
syntax error: unexpected [, expecting field name or embedded type  
```

### 2.3 在其他包使用

如果一个含有`unexported` field的struct需要在包外使用，需要在包外初始化该对象，那么可以通过`package.NewT()`, 方法来作为该对象的一个构造器。即使内部嵌套类型有一个无法导出的域，也需要使用该方式。

```go
// post/website.go
package post

import "fmt"

// define the websites struct which contains posts
type WebSite struct {
  Posts []Post
}

// display the website msssage
func (w WebSite) Display() {
  fmt.Println("The content of the website:\n")
  for _, po := range w.Posts {
    po.Detail()
    fmt.Println()
  }
}

```

```go
// post/post.go
package post

/* define the post struct */
import (
  "fmt"
)

// author of the post
type author struct {
  firstName, lastName string
  bio string
}

// get the author full name
func (a author) AuthorName() string {
  return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func Newauthor(firstName, lastName, bio string) author {
  return author{firstName, lastName, bio}
}

// struct of post, `author` field promoted
// so Post type can use `AuthorName` directly
type Post struct {
  Title string
  Content string
  author
}

func NewPost(title, content string, a author) Post {
  return Post{title, content, a}
}

func (p Post) Detail() {
  fmt.Println("Post:", p.Title)
  fmt.Println("Content:", p.Content)
  fmt.Println("Author:", p.AuthorName())
  fmt.Println("Bio:", p.bio)
}
```

## 3. 多态 Polymorphism

在golang中使用接口来实现多态。通过定义同一类型的接口，接口可以被隐式的实现。需要使用这个接口中的类型可以实现接口中的所有方法，然后使用接口变量指向这些类型变量。实现多态操作。使用接口的最大好处在于提高了代码的可复用性，如果有新的类型需要加入到代码中，不需要更改其他部分，只需要添加该类型以及接口的实现。

**一个接口类型的变量，可以指向任何一个实现了该接口的值。使用该特性来实现多态。**

```go
package main

import (  
    "fmt"
)

type Income interface {  
    calculate() int
    source() string
}

type FixedBilling struct {  
    projectName string
    biddedAmount int
}

type TimeAndMaterial struct {  
    projectName string
    noOfHours  int
    hourlyRate int
}

func (fb FixedBilling) calculate() int {  
    return fb.biddedAmount
}

func (fb FixedBilling) source() string {  
    return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {  
    return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {  
    return tm.projectName
}

func calculateNetIncome(ic []Income) {  
    var netincome int = 0
    for _, income := range ic {
        fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
        netincome += income.calculate()
    }
    fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {  
    project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
    project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
    project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
    incomeStreams := []Income{project1, project2, project3}
    calculateNetIncome(incomeStreams)
}
```

**注意使用接口时，如果是指针类型实现了接口的方法，那么指针类型才可以赋值给该接口类型。不存在指针的隐式转换。只有在方法中存在指针的隐式转换。**