# API-BASE

一个正在不断完善中的、由 Go 语言编写的、工程化的 HTTP 接口服务项目。

## 编写目的

详细地阐述一个基于 Go 的通用接口工程项目的标准化构建流程。

## 关键字

- Router [Gin](https://github.com/gin-gonic/gin)
- Middleware [Gin](https://github.com/gin-gonic/gin)
- Database [GORM](https://github.com/jinzhu/gorm)
- CI/CD [GitHub Actions](https://github.com/actions)

[//]: # (- Document [Swagger]&#40;https://swagger.io/&#41; 生成)

[//]: # (- Config [Viper]&#40;https://github.com/spf13/viper&#41;)

[//]: # (- Auth [JWT]&#40;https://jwt.io/&#41;)

[//]: # (- Validator [validator]&#40;https://github.com/go-playground/validator&#41;)

[//]: # (- Cron [cron]&#40;https://github.com/robfig/cron&#41;)

[//]: # (- Test [GoConvey]&#40;http://goconvey.co/&#41;)

[//]: # (- Lint [GolangCI-lint]&#40;https://golangci.com/&#41;)

## 项目目录

```
|-- app
|-- cmd     # 脚手架目录（待完善）
|-- config      # 配置文件
|   |-- local   # 开发环境配置
|   |   |-- app.yaml        # app 配置
|   |   |-- database.yaml   # 数据库配置
|   |-- prod    # 部署环境配置
|   |   |-- app.yaml
|   |   |-- database.yaml
|   |-- test    # 测试环境配置
|   |   |-- app.yaml
|   |   |-- database.yaml
|   |-- config.go   # 配置项的类型声明
|-- docs    # 相关文档（待完善）
|-- internal    # 业务目录
|   |-- api     # http 接口
|   |-- ecode   # 自定义业务错误码
|   |-- model   # 数据库 model
|   |-- repo    # 数据访问层  
|   |-- router  # Gin 以及业务路由
|   |-- service # 业务逻辑层
|-- pkg     # 公用的 package
|   |-- config  # 配置文件读取功能封装
|   |-- errno   # 公用错误码以及自定义错误方法封装
|   |-- response    # http 请求返回方法封装
|   |-- middleware  # 公用的中间件
|   |-- snowflake   # 雪花ID生成工具
|-- main.go     # 项目入口文件
```

## config 目录

存储配置文件，默认从 config/{ENV} 加载配置。

如果是多环境可以设定不同的配置目录，比如：

- config/local 本地开发环境
- config/test 测试环境
- config/staging 预发布环境
- config/prod 线上环境

## internal

### model 目录

`model` 层，或者叫 `entity` 层、实体层。

用于存放实体类，为数据库表的映射。

其中，`mysql.go` 定义了 `MySQL` 的连接方法。

相关约定：

- 默认使用 `MySQL` 数据库。

- 使用 `GORM` 作为 ORM 库。

- 声明实体类的结构体符合 `GORM` 的约定规范，一个表中需要包含的三大字段：
    - 主键(id)
    - 创建时间(created_at)
    - 更新时间(updated_at)

### repo 目录

仓库层，处于 `service` 层 与 `model` 层之间。

负责对数据库等数据的访问，对上层屏蔽数据访问细节。

更换、升级 ORM 引擎时，不影响业务逻辑。

单元测试时，用 `Mock` 对象代替实际的数据库存取，可以成倍地提高测试用例运行速度。

职责：

- DB 访问逻辑
- DB 的拆库分表逻辑
- DB 的缓存读写逻辑
- HTTP 接口调用逻辑

相关约定：

- 禁止使用连表查询

#### repo.go

该文件中包括了仓库类的类型声明，构造函数与所有类方法的声明。

仓库类的类型声明，其中包含了数据库 `ORM`，例如：

```
type repository struct {
    orm *gorm.DB
    db  *sql.DB
}
```

仓库类方法的声明，例如：

```
type Repository interface {
    GetFoo(ctx context.Context) (*model.Foo, error)
    CreateFoo(ctx context.Context, foo *model.Foo) (string, error)
    UpdateFoo(ctx context.Context, foo *model.Foo) error
    DeleteFoo(ctx context.Context, id string) error
}
```

仓库类构造函数，例如：

```
func New(db *gorm.DB) Repository {
    return &repository{
        orm: db,
    }
}
```

#### *_repo.go

每个业务分一个文件，其中每一个 `*_repo.go` 文件对应一个表操作。

在对应的文件中将 `repo.go` 中定义的仓库类方法逐一实现。

[//]: # (## 单元测试)

[//]: # (关于数据库的单元测试可以用到的几个库：)

[//]: # (- go-sqlmock https://github.com/DATA-DOG/go-sqlmock 主要用来和数据库的交互操作:增删改)

[//]: # (- GoMock https://github.com/golang/mock)

### service 目录

业务逻辑层，处于 `api` 层和 `repo` 层之间。

职责：

- 接受 `api` 层的调用，接收其传来的参数。
- 通过整合 `repo` 层的方法，获取数据并处理，完成一个相对完整的业务功能。
- 负责第三方接口的请求

#### service.go

该文件中包含了总服务类的定义，总服务类的类型声明，构造函数与类方法的声明与实现。

其中，在总服务类方法的实现中通过调用各模块服务的构造函数来进行初始化。

总服务类的声明，服务类初始化的入口：

```
var Svc Service
```

总服务类的类型声明：

```
type service struct {
    repo repo.Repository
}
```

总服务类方法的声明与实现：

```
type Service interface {
    Foo() FooService
    Bar() BarService
}

func (s *service) Foo() FooService {
    return newFooSvc(s)
}

func (s *service) Bar() BarService {
    return newBarSvc(s)
}

```

总服务类的构造函数：

```
func New(repo repo.Repository) Service {
    return &service{
        repo: repo,
    }
}
```

#### *_service.go

该文件中包括了各模块服务类的类型声明，构造函数与类方法的声明与实现。

模块服务类的类型声明，包含了仓库实例，例如：

```
type fooService struct {
    repo repo.Repository
}
```

模块服务类方法的声明与实现，例如：

```
type FooService interface {
    GetFoo(ctx context.Context, id string) (*model.Foo, error)
    AddFoo(ctx context.Context, foo *model.Foo) (string, error)
    UpdateFoo(ctx context.Context, foo *model.Foo) error
    DeleteFoo(ctx context.Context, id string) error
}

func (ps *fooService) GetFoo(ctx context.Context, id string) (*model.Foo, error) {
    // ...
}

func (ps *fooService) AddFoo(ctx context.Context, foo *model.Foo) (string, error) {
    // ...
}

func (ps *fooService) UpdateFoo(ctx context.Context, foo *model.Foo) error {
    // ...
}

func (ps *fooService) DeleteFoo(ctx context.Context, id string) error {
    // ...
}
```

模块服务类的构造函数，例如：

```
func newFooSvc(svc *service) *fooService {
    return &fooService{repo: svc.repo}
}
```

### api 目录

HTTP 接口层，将接口服务暴露给外部，供外部调用，处于 `service` 层之上。

由 `Gin Router` 统一组织、分配路由。

由类似 `v1` 的目录组织来表示接口版本。

每个接口方法作为一个单独的文件。

职责：

- 响应 HTTP 请求
- 请求参数的验证
- `service` 层的调用
- 请求响应的发送

### router 目录

`Gin router` 服务的初始化。

内容包括：

- 接口总 `service` 的初始化
- `Gin router` 初始化
- 将中间件的加载至 `router` 实例
- 将一些公用的异常处理方法的加载至 `router` 实例
- 将 `HealthCheck` 方法的加载至 `router` 实例
- 使用 `Gin.Group` 将各版本的接口方法组织后统一加载至 `router` 实例
- `router.Run` 方法调用，启动接口服务

### ecode 目录

自定义业务错误码，根据业务实际定义的统一错误码，通用性较弱。

可以根据模块按文件进行定义。

以 `ecode.开头`，例如：

```
ecode.ErrUserNotFound
```

错误码的结构说明可参考 [公共系统错误码](#错误代码说明)

## pkg

### config

配置文件读取功能封装。

包括：

- 配置文件类的类型声明
- 配置文件类的构造函数
- 有关配置文件类的其他实用方法

### errno

公共系统错误码。

#### 错误代码说明

| 1      | 00     | 02     |
|:-------|:-------|:-------|
| 错误级别代码 | 服务模块代码 | 具体错误代码 |

- 服务级别错误：
    - 1 为系统级错误；
    - 2 为服务级错误，通常是由用户非法操作引起，返回业务错误码。
- `code = 0` 说明是正确返回，`code > 0` 说明是错误返回
- 上述说明同样适用于业务错误码。

[//]: # (## Reference)

[//]: # (- [gRPC错误处理]&#40;https://mp.weixin.qq.com/s/ghJiTvJxYzLKTFs5gZga5w&#41;)

### middleware

下面是一些常用的中间件：

- gin-jwt：用于 Gin 框架的 JWT 中间件
- gin-sessions：基于 MongoDB 和 MySQL 的会话中间件
- gin-location：用于公开服务器主机名和方案的中间件
- gin-nice-recovery：异常错误恢复中间件，让您构建更好的用户体验
- gin-limit：限制同时请求，可以帮助解决高流量负载
- gin-oauth2：用于处理 OAuth2
- gin-template：简单易用的 Gin 框架 HTML/模板
- gin-redis-ip-limiter：基于 IP 地址的请求限制器
- gin-access-limit：通过指定允许的源 CIDR 表示法来访问控制中间件
- gin-session：Gin 的会话中间件
- gin-stats：轻量级且有用的请求指标中间件
- gin-session-middleware：一个高效，安全且易于使用的 Go 会话库
- ginception：漂亮的异常页面
- gin-inspector：用于调查 HTTP 请求的 Gin 中间件
- RestGate：REST API 端点的安全身份验证

参考：https://github.com/chenjiandongx/ginprom

### response

定义了公共的 `Response` 类型，作为全局统一的接口返回结构。

### snowflake

雪花 ID 生成工具

### utils

其他一些实用工具。

## scripts

脚本目录，存放用于执行各种构建，安装，分析等操作的脚本。

[//]: # (Makefile 中执行的一些脚本可以放到这里，让 Makefile 变得更小巧、简单。)

[//]: # ()

[//]: # (脚本文件)

[//]: # ( - admin.sh             # 进程的start|stop|status|restart控制文件)

[//]: # ( - wrktest.sh           # API 性能测试脚本)

[//]: # ( )

[//]: # (## Reference)

[//]: # (- 压测工具 https://github.com/tsenart/vegeta)

## main.go

系统的总入口文件，职责如下：

- 读取配置文件
- 初始化时区
- 初始化日志
- 初始化数据库连接
- 初始化接口服务
- 启动 HTTP 服务

> 参考
>
> 本项目主要参考了 [Eagle](https://github.com/go-eagle/eagle) 框架，并根据实际需求进行了改造。感谢该框架提供的思路与学习方向。
>
> API 设计参考：[Google Cloud API 设计指南](https://cloud.google.com/apis/design?hl=zh-cn)
>
> 错误码设计参考：[新浪开放平台 Error code](http://open.weibo.com/wiki/Error_code)
>
> 开发规范参考：[Uber Go 语言编码规范](https://github.com/uber-go/guide/blob/master/style.md)
> 
