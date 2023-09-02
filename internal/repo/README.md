# repo

数据访问层，处于 `Service` 层 与 DB 之间。

负责对数据库等数据的访问，对上层屏蔽数据访问细节。  

[//]: # (更换、升级 ORM 引擎，不影响业务逻辑。)

[//]: # (能提高测试效率，单元测试时，用Mock对象代替实际的数据库存取，可以成倍地提高测试用例运行速度。)

职责：
- DB 访问逻辑
- DB 的拆库分表逻辑
- DB 的缓存读写逻辑
- HTTP 接口调用逻辑

> Tips: 如果是返回的列表，尽量返回map，方便service使用。

相关约定：
- 禁止使用连表查询


## repo.go

所有 repo 接口的定义，

```
type Repository interface {
	// BaseUser 
	...
	// Follow
    ...
	// Stat
	...
}
```

## 各分目录

每个业务分一个目录，其中每一个 `*_repo.go` 文件对应一个表操作。

如，在 `/portlet` 目录下：
- Portlet 基础表：portlet_base_repo.go
- Portlet 角色表：portlet_role_repo.go
- Portlet 用户表：portlet_user_repo.go

[//]: # (## 单元测试)

[//]: # ()
[//]: # (关于数据库的单元测试可以用到的几个库：)

[//]: # (- go-sqlmock https://github.com/DATA-DOG/go-sqlmock 主要用来和数据库的交互操作:增删改)

[//]: # (- GoMock https://github.com/golang/mock)