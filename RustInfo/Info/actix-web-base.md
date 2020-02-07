# Actix-web 基本使用

Actix-web  是基于`Actix` actor 异步并发框架以及`tokio`异步IO系统的web高性能框架，具有类型安全，易于扩展，特性丰富等优势。 是一个轻量且实用的web框架。Actix-web 不仅可以用作微服务部署于nginx等HTTP server之后，而且可以作为单独的HTTP server使用，支持HTTP 1， HTTP 2以及SSL/TSL。 Actix-web 需要在rust 1.39之后运行.

## 1. APP

### 1.1 启动

`actix-web`提供了各种原型构建server以及app。提供了routing，middleware， pre-processing request，post-processing response。一个`HttpServer`实例接受一个`web::app`作为参数，`server`用以绑定端口，运行异步服务，`app`用以注册资源路由、中间件、存储应用状态（可以被同一个命名域下`handlers`共享）。应用的`scope`扮演了所有路由的命名域，在此`scope`下的路由规则具有相同的资源前缀。

使用`actix-web`以及`actix-rt`，在`Cargo.toml`添加对应依赖.

```rust
use actix_web::{web, App, HttpResponse, HttpServer};

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(
        || App::new()
            .service(web::resource("/").to(|| HttpResponse::Ok())))
        .bind("127.0.0.1:59090")?
        .run()
        .await
}
```

使用`await`延迟执行，直到`Future`处于ready状态。

### 1.2 应用状态

应用状态可以被处于同一个`scope`下的所有router共享，使用`web::Data<T>`进行存储, 应用状态对于中间件来讲也是可用的。`HttpServer`接受一个app 工厂作为参数而不是实例，`HttpServer`为每一个线程构建一个`app`实例，因此应用数据一定被多次创建。**如果需要在多个线程间共享数据，必须使用共享对象，例如`Arc`**. 内部的`Data`类型使用了`Arc`，所以`Data`可以在app factory外创建，并且使用`clone`在`App::app_data()`方法调用中存储。

```rust
use std::cell::Cell;
use actix_web::{web, App, HttpResponse, Responder};

struct MyData {
    counter: Cell<usize>,
}

async fn index(data: web::Data<MyData>) -> impl Responder {
    data.counter.set(data.counter.get() + 1);
    HttpResponse::Ok()
}

let app = App::new()
    .data(MyData{ counter: Cell::new(0) })
    .service(
        web::resource("/index.html").route(
            web::get().to(index)));
```

`App::app_data()` 方法可以存放应用全局数据，可以使用`HttpRequest::app_data`在运行时获取 。

### 1.3 命名域

使用`web::scope`方法设定具有相同前缀的routers。

```rust
#[actix_rt::main]
async fn main() {
    App::new()
        .service(
            web::scope("/users")
                .route("/show", web::get().to(show_users)));
}
```

### 1.4 route guard

`guard`通过设定相应的条件，例如`method`，`headers`等，过滤不符合条件的请求。

```rust
 App::new()
            .app_data(data.clone()) // register data
            .configure(config)
            .service(
                web::scope("/api")
                // use guard to set route or scope filter
                .guard(guard::Header("Token-access", "justatoken"))
                .service(index)
                .service(sim_obj)
                // 由于Scope并没有指定唯一的资源实体，所以需要path参数指定
                // 唯一URI，然后指定相应的router
                .route("/hello", web::get().to(hello))
                .service(name)
            )
```

### 1.5 服务配置

为了提高`App`以及`web::scope`的可重用性，提供了`configure`方法,可以将配置部分移动至不同的模块或者不同的库中.

```rust
use actix_web::{web, App, HttpResponse, HttpServer};

// this function could be located in different module
fn scoped_config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::resource("/test")
            .route(web::get().to(|| HttpResponse::Ok().body("test")))
            .route(web::head().to(|| HttpResponse::MethodNotAllowed())),
    );
}

// this function could be located in different module
fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::resource("/app")
            .route(web::get().to(|| HttpResponse::Ok().body("app")))
            .route(web::head().to(|| HttpResponse::MethodNotAllowed())),
    );
}

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .configure(config)
            .service(web::scope("/api").configure(scoped_config))
            .route("/", web::get().to(|| HttpResponse::Ok().body("/")))
    })
    .bind("127.0.0.1:8088")?
    .run()
    .await
}
```

响应请求结果:

```http
/         -> "/"
/app      -> "app"
/api/test -> "test"
```

## 2. Router

对于`app`或者`web::scope`对象配置路由时可以使用`route`或者`service`方法：

`route`方法接受两个参数，第一个指定url路径，第二个指定了一个`Router`对象，`Router`可以使用`method`方法指定请求的方法，使用`to`方法绑定到具体的`handler`。为了方便使用可以直接用`web::get, web::post, web::patch`等方法。其内部使用了`Router::method(actix_http::Method::Get)`等。

```rust
use actix_web::{web, App, HttpResponse};

let app = App::new().service(
    web::resource("/{project_id}")
       .route(web::get().to(|| HttpResponse::Ok()))
);
```

也可以使用`pub fn service<T: IntoPattern>(path: T) -> WebService`：
配合handler宏调用定义路由。

```rust
#[post("/users")]
async fn create_users(info: web::Json<User>) -> impl Responder {
    info.into_inner()
}

let app = App::new().service(
                web::scope("/app")
                    // use a fn_guard to filter
                    .guard(guard::fn_guard(
                        |req| req.headers()
                            .contains_key("content-type")
                    ))
                    .service(create_users)
            )
```

## 2. 增量更新

进行web开发时为了实现快速开发以及测试，可以借助`systemfd`， `cargo-watch` 进行自动加载。`systemfd`会开启一个`socket`连接，并传递给`cargo-watch`，`cargo-watch`负责监控代码变化，并且进行实时编译，并且将服务暴露于`systemfd`提供的`socket`端口下。

+ 安装插件

```shell
cargo install systemfd cargo-watch
```

```toml
[dependencies]
listenfd = "0.3"
```

+ 代码增添

  ```rust
  use actix_web::{web, App, HttpRequest, HttpServer, Responder};
  use listenfd::ListenFd;
  
  async fn index(_req: HttpRequest) -> impl Responder {
      "Hello World!"
  }
  
  #[actix_rt::main]
  async fn main() -> std::io::Result<()> {
      let mut listenfd = ListenFd::from_env();
      let mut server = HttpServer::new(|| App::new().route("/", web::get().to(index)));
  
      server = if let Some(l) = listenfd.take_tcp_listener(0).unwrap() {
          server.listen(l)?
      } else {
          server.bind("127.0.0.1:3000")?
      };
  
      server.run().await
  }
  ```

+ 服务启动

```shell
systemfd --no-pid -s http::3000 -- cargo watch -x run
```

