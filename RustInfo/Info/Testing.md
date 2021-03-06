# 测试

[TOC]

## 1. 编写测试的基本操作

 Rust中使用函数进行代码的测试，用以验证非测试代码是否按照期望的方式运行，一般具有以下操作：

+ 设置所需要的数据或者状态
+ 运行测试代码
+ 断言结果是否为所期望的

### 1.1 test 属性的注释函数

Rust 中测试函数是一个带有`test`属性注释的函数。属性attribute是关于Rust代码片段的元数据，`derive`就是一个典型的例子。为了将一个函数设置为测试函数，可以在函数名前加上一个`#[test]`.当执行`cargo test` 时，Rust会构建一个测试程序来调用所有标记了该属性注释的函数。

对于一个成本的测试样例，得到的运行结果：

```rust
running 7 tests
test tests::it_works ... ok
test tests::rectangle_large_can_hold_small ... ok
test tests::greeting_not_done ... FAILED
test tests::length_negative ... ok
test tests::rectangle_small_can_not_hold_large ... ok
test tests::return_result ... FAILED
test tests::width_negative ... ok

failures:

---- tests::greeting_not_done stdout ----
thread 'tests::greeting_not_done' panicked at 'Greeting did not contain the name value:Hello, Cargo', src/lib.rs:40:9
note: Run with `RUST_BACKTRACE=1` for a backtrace.

---- tests::return_result stdout ----
Error: "two add two does not equal three"
thread 'tests::return_result' panicked at 'assertion failed: `(left == right)`
  left: `1`,
 right: `0`: the test returned a termination value with a non-zero status code (1) which indicates a failure', src/libtest/lib.rs:337:5


failures:
    tests::greeting_not_done
    tests::return_result

test result: FAILED. 5 passed; 2 failed; 0 ignored; 0 measured; 0 filtered out
```

输出了所有的失败样例，以及测试结果总结，

通常在一个测试函数中，使用`assert！`类的宏进行测试的第三项工作，该宏接受一个`bool`类型的值，通常使用的宏还有`assert_eq!`, `assert_ne!`,接受两个类型相同的值，断言相等或者不等。**需要注意的是，在一些语言和测试框架中，断言两个值相等的函数的参数叫做 `expected` 和 `actual`，而且指定参数的顺序是很关键的。然而在 Rust 中，他们则叫做 `left` 和 `right`，同时指定期望的值和被测试代码产生的值的顺序并不重要。这个测试中的断言也可以写成 `assert_eq!(add_two(2), 4)`，这时失败信息会变成 `assertion failed: `(left == right)`` 其中 `left` 是 `5` 而 `right` 是 `4`。**

这两个宏底层实现使用了`==  !=` ， 同时如果断言失败会使用`Debug`模式输出错误信息。所以传入的类型必须是实现了`PartialEq`以及`Debug` trait 的类型。又因为这两个trait都是派生trait，通常可以在结构体上加`#[derive(PartialEq + Debug)]`属性注释。

### 1.2 自定义失败信息

可以向`assert！`等宏中添加一个可选的默认失败信息，在测试失败时将这些信息打印出来。获取更加详细的测试信息。

例如，对于一个尚未完成的函数进行测试，由于可能`Hello`可变，只对参数进行测试：

```rust
# fn main() {}
pub fn greeting(name: &str) -> String {
    format!("Hello {}!", name)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn greeting_contains_name() {
        let result = greeting("Carol");
        assert!(result.contains("Carol"));
    }
}
```

如果测试失败，显示的结果只有失败的函数以及行号：

```shell
running 1 test
test tests::greeting_contains_name ... FAILED

failures:

---- tests::greeting_contains_name stdout ----
        thread 'tests::greeting_contains_name' panicked at 'assertion failed:
result.contains("Carol")', src/lib.rs:12:8
note: Run with `RUST_BACKTRACE=1` for a backtrace.

failures:
    tests::greeting_contains_name
```

可以添加更多信息：

```rust
#[test]
fn greeting_contains_name() {
    let result = greeting("Carol");
    assert!(
        result.contains("Carol"),
        "Greeting did not contain name, value was `{}`", result
    );
}
```

### 1.3 panic 测试

对于代码中的panic进行测试时，可以使用`should_panic`属性注释进行操作：

```rust
#[test]
#[should_panic(expected = "The incorrect value of rectangle: 12.1, -5")]
fn width_negative() {
        Rectangle::new(12.1, -5.);
}

impl Rectangle {
        pub fn new(length: f64, width: f64) -> Rectangle {
            if length <= 0. || width <= 0. {
                panic!("The incorrect value of rectangle: {}, {}", length, width);
            }
            Rectangle {
                length,
                width,
            }
        }
}

```

但是使用这种方式有一个很大的问题，如果是程序的其他部分panic那么该测试样例也是可以通过的，是的测试结果十分含糊不清，所以一般需要加上一个`expeted`参数，从而获取所期望的panic信息。

### 1.4 用`Result<T, E>`进行测试

可以在测试函数返回一个`Result<T, E>`枚举类型，可以达到同样的效果：

```rust
mod tests {
    #[test]
    fn it_works() -> Result<(), String> {
        if 2 + 2 == 4 {
            Ok(())
        } else {
            Err(String::from("two plus two does not equal four"))
        }
    }
}
```

这里我们将 `it_works` 改为返回 Result。同时在函数体中，在成功时返回 `Ok(())` 而不是 `assert_eq!`，而失败时返回带有 `String` 的 `Err`。跟之前一样，这个测试可能成功或失败，不过不再通过 panic，可以通过 `Result<T, E>` 来判断结果。为此不能在对这些函数使用 `#[should_panic]`；而是应该返回 `Err`！  

## 2. 测试控制

`cargo test`命令会生成可执行的二进制文件，默认进行的是并行测试，并且截获所有的输出，阻止其输出到console，从而获得更为清晰地测试结果。可以将一部分命令行参数传递给`cargo test`，而另一部分传递给生成的二进制文件，先列出`cargo test`的参数，然后是二进制文件参数。

### 2.1 并行或者串行执行测试

并执行测试时必须保证测试程序间不会生竞争条件，所以测试间不可以相互依赖，或者依赖共享的状态。依赖共享的环境。为了避免测试相互干扰，可以指定单线程运行：

```shell
cargo test -- -- test-thread=1 
```

### 2.2 显示函数输出

```shell
cargo test -- -- nocapture
```

### 2.3 运行单个测试

```rust
cargo test func_name
```

可以通过指定一个测试函数或者测试模块，运行一部分测试，节省测试时间。这种方式称为`filter`

### 2.4 忽略某些测试

对于一些较为耗时的测试，可以添加`#[ignore]`属性注释进行忽略，从而加快测试进度。

## 3. 测试的组织结构

本章一开始就提到，测试是一个复杂的概念，而且不同的开发者也采用不同的技术和组织。Rust 社区倾向于根据测试的两个主要分类来考虑问题：**单元测试**（*unit tests*）与 **集成测试**（*integration tests*）。单元测试倾向于更小而更集中，在隔离的环境中一次测试一个模块，或者是测试私有接口。而集成测试对于你的库来说则完全是外部的。它们与其他外部代码一样，通过相同的方式使用你的代码，只测试公有接口而且每个测试都有可能会测试多个模块。

为了保证你的库能够按照你的预期运行，从独立和整体的角度编写这两类测试都是非常重要的。

### 3.1 单元测试

单元测试的主要目的是在与其他部分隔离的环境中进行代码的测试。以便于快速而准确地确定某个部分的代码是否符合预期。单元测试与需要测试的代码都存放在`src`目录下，规范是包含测试函数的`tests`模块，需要使用`#cfg(test)`属性标注。

通过使用`#cfg(test)`保证了测试代码只在`cargo test`时才编译运行，在`cargo build`时不需要这么做。可以在构建库时节省时间。同时由于单元测试代码和库文件在同一个文件夹中，所以需要这个属性进行标示，指定他们不可以包含在库文件中。

### 3.2 集成测试

在 Rust 中，集成测试对于所需要测试的库是完全外部的。同其他使用库的代码一样使用库文件，也就是说只可以调用库中的`pub`API。集成测试的目的是为了保证多个部分是否可以协同工作。集成测试在`tests`文件夹内进行。

`tests`目录与`scr`目录同级，不需要加`#cfg(test)`属性 `tests` 文件夹在 Cargo 中是一个特殊的文件夹， Cargo 只会在运行 `cargo test` 时编译这个目录中的文件。

```shell
  Compiling adder v0.1.0 (/home/inno/Learning-notes/RustInfo/projects/adder)
    Finished dev [unoptimized + debuginfo] target(s) in 0.23s
     Running target/debug/deps/adder-70257eb04d3ac84f

running 7 tests
test tests::greeting_not_done ... ok
test tests::it_works ... ok
test tests::rectangle_large_can_hold_small ... ok
test tests::rectangle_small_can_not_hold_large ... ok
test tests::length_negative ... ok
test tests::return_result ... ok
test tests::width_negative ... ok

test result: ok. 7 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out

     Running target/debug/deps/integration_test-b8e9d862b2a5bae4

running 1 test
test it_adds_two ... ok

test result: ok. 1 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out

   Doc-tests adder

running 0 tests

test result: ok. 0 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out

```

现在有了三个部分的输出：单元测试、集成测试和文档测试。仍然可以通过指定测试函数的名称作为 `cargo test` 的参数来运行特定集成测试。也可以使用 `cargo test` 的 `--test` 后跟文件的名称来运行某个特定集成测试文件中的所有测试：