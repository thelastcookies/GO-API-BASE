# WEB-SERVICE

## Intro

一个正在不断完善中的、由 Go 语言编写的、工程化的 HTTP 接口服务项目。

## Project Structure

```
|-- app
|-- cmd                         # 脚手架目录（待完善）
|-- config                      # 配置文件
|-- docs                        # 相关文档（待完善）
|-- internal                    # 业务目录
    |-- api                     # http 接口
    |-- ecode                   # 自定义业务错误码
    |-- model                   # 数据库 model
    |-- repo                    # 数据访问层  
    |-- router                  # Gin 以及业务路由
    |-- service                 # 业务逻辑层
|-- pkg                         # 公用的 package
    |-- config                  # 配置文件读取功能封装
    |-- errno                   # 公用错误码以及自定义错误方法封装
    |-- response                # http 请求返回方法封装
    |-- middleware              # 公用的中间件
    |-- snowflake               # 雪花ID生成工具
|-- main.go                     # 项目入口文件
```

## Current Features

- Router [Gin](https://github.com/gin-gonic/gin)
- Middleware [Gin](https://github.com/gin-gonic/gin)
- Database [GORM](https://github.com/jinzhu/gorm)

[//]: # (- Document [Swagger]&#40;https://swagger.io/&#41; 生成)
[//]: # (- Config [Viper]&#40;https://github.com/spf13/viper&#41;)
[//]: # (- Auth [JWT]&#40;https://jwt.io/&#41;)
[//]: # (- Validator [validator]&#40;https://github.com/go-playground/validator&#41;)
[//]: # (- Cron [cron]&#40;https://github.com/robfig/cron&#41;)
[//]: # (- Test [GoConvey]&#40;http://goconvey.co/&#41;)
[//]: # (- CI/CD [GitHub Actions]&#40;https://github.com/actions&#41;)
[//]: # (- Lint [GolangCI-lint]&#40;https://golangci.com/&#41;)

## 


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
